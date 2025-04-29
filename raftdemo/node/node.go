package node

import (
	"raft/raftdemo/consensus"
	"raft/raftdemo/network"
	"raft/raftdemo/raftlog"
	"raft/raftdemo/statemachine"
)

type Node struct {
	ID        string                    // ID del nodo
	Raft      consensus.Raft            // interfaccia per il protocollo Raft
	Transport network.Transport         // interfaccia per la rete
	Log       raftlog.Log               // interfaccia per il log
	State     statemachine.StateMachine // interfaccia per la macchina a stati
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
	n.Raft.Tick() // Avvia il tick del protocollo Raft (heartbeat)
}
