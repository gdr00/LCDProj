package main

import "raft/raftdemo/test"

func main() {
	// Crea un cluster di nodi
	cluster := test.NewCluster(5)

	// Avvia il cluster
	cluster.Start()

	// Esegui test di consenso
	test.RunConsensusTests(cluster)

	// Ferma il cluster
	cluster.Stop()
}
