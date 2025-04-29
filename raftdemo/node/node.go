package node

import (
	"raft/raftdemo/consensus"
	"raft/raftdemo/network"
	"raft/raftdemo/raftlog"
	"raft/raftdemo/statemachine"
)

type Node struct {
	ID        string
	Raft      consensus.Raft
	Transport network.Transport
	Log       raftlog.Log
	State     statemachine.StateMachine
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
