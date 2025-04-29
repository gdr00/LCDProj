package main

import (
	"raft/raftdemo/network"
	"raft/raftdemo/node"
)

func main() {
	transport := network.Newtransport()

	n1 := node.NewNode("node1", transport)
	n2 := node.NewNode("node2", transport)
	n3 := node.NewNode("node3", transport)

	go n1.Start()
	go n2.Start()
	go n3.Start()

	select {} // blocca per sempre
}
