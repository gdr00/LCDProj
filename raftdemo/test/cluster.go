package test

import (
	"raft/raftdemo/node"
	"strconv"
)

// Cluster rappresenta un cluster di nodi Raft
type Cluster struct {
	nodes []*node.Node // Array di nodi
}

func NewCluster(numNodes int) *Cluster {
	nodes := make([]*node.Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = node.NewNode(strconv.Itoa(i), nil) // Passa il transport appropriato
	}
	return &Cluster{nodes: nodes}
}

func (c *Cluster) Start() {
	for _, n := range c.nodes {
		n.Start() // Avvia ogni nodo
	}
}

func (c *Cluster) Stop() {
	for _, n := range c.nodes {
		n.Stop() // Ferma ogni nodo
	}
}

func RunConsensusTests(cluster *Cluster) {
	// Qui puoi implementare i test di consenso
	// Ad esempio, inviare comandi e verificare lo stato dei nodi
	for _, n := range cluster.nodes {
		n.Raft.Submit([]byte("set 5")) // Invia un comando di test
	}
}
