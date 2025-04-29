package consensus

type Raft interface {
	Tick()                                // chiamato ogni tot millisecondi
	HandleMessage(msg any)                // riceve messaggi da altri nodi
	Submit(command []byte) (string, bool) // chiamato dal client tornando un ID o ok
}
