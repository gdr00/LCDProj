package consensus

type Raft interface {
	Tick()                 // chiamato ogni tot millisecondi
	HandleMessage(msg any) // riceve messaggi da altri nodi
	Submit(command []byte) // chiamato dal client
}
