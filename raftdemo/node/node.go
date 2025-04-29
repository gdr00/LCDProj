package node

import (
	"raft/raftdemo/consensus"
	"raft/raftdemo/logger"
	"raft/raftdemo/network"
)

type Node struct {
	ID        string            // ID del nodo
	Raft      consensus.Raft    // interfaccia per il protocollo Raft
	Transport network.Transport // interfaccia per la rete
	Log       logger.Log        // interfaccia per il log
	State     consensus.State   // stato del nodo (Leader, Follower, Candidate)
}

func NewNode(id string, transport network.Transport) *Node {
	// Qui istanzierai tutte le dipendenze concrete
	return &Node{
		ID:        id,
		Transport: transport,
		// Raft, Log, State: inizializzazioni vere da fare dopo
	}
}

func (n *Node) Start() {
	// Avvia tick periodico + registra handler messaggi
}

func (n *Node) Stop() {
	// Ferma il nodo e chiude le connessioni
	n.Transport.Close() // Chiude le connessioni di rete
}
