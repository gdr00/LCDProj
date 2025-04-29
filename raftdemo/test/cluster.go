package test

import (
	"log"
	"raft/raftdemo/node"
	"strconv"
)

// Cluster rappresenta un cluster di nodi Raft
type Cluster struct {
	nodes []*node.Node // Array di nodi
}

func NewCluster(numNodes int) *Cluster {
	// Crea il transport per la comunicazione tra i nodi

	// Crea un array di nodi Raft
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
	cluster.nodes[1].Raft.Submit([]byte("set 5")) // Invia un comando di test
	log.Println("test scrittura su nodo 1 set 5\n")
}
