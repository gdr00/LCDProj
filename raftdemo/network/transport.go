package network

import (
	"fmt"
	"sync"
)

type Transport interface {
	Send(msg Message) error
	Receive() (<-chan Message, error)
	Broadcast(msg Message) error
	Close() error
}

var _ Transport = (*ChanTransport)(nil) // verifica che ChanTransport implementi l'interfaccia Transport

type ChanTransport struct {
	peers map[string]chan<- Message // canali per ogni peer per l'invio di messaggi
	ch    chan Message              // canale per ricevere messaggi
	mu    sync.RWMutex
}

func NewChanTransport(peers []string) *ChanTransport { //TODO: fix NewChanTransport per condividere i canali
	// Crea un nuovo ChanTransport con i peer specificati
	peersMap := make(map[string]chan<- Message, len(peers))
	for _, peer := range peers {
		peersMap[peer] = make(chan Message, 20) // canale per ogni peer con buffer di 20
	}

	return &ChanTransport{
		peers: peersMap,               // mappa dei peer e i loro canali
		ch:    make(chan Message, 20), // canale per ricevere messaggi buffer di 20?
	}
}

func (t *ChanTransport) RegisterPeer(id string, ch chan Message) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.peers[id] = ch
}

func (t *ChanTransport) Send(msg Message) error {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if ch, ok := t.peers[msg.To]; ok {
		select {
		case ch <- msg:
			return nil
		default:
			return fmt.Errorf("channel for peer %s is full", msg.To)
		}
	} else {
		return fmt.Errorf("peer %s not found", msg.To)
	}
}

func (t *ChanTransport) Receive() (<-chan Message, error) {
	return t.ch, nil
}

func (t *ChanTransport) Broadcast(msg Message) error { //TODO: implementare broadcast
	t.mu.RLock()
	defer t.mu.RUnlock()

	for id, ch := range t.peers {
		select {
		case ch <- msg:
		default:
			return fmt.Errorf("channel for peer %s is full", id)
		}
	}
	return nil
}

func (t *ChanTransport) Close() error {
	close(t.ch)
	return nil
}
