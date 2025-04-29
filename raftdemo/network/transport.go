package network

const (
	MsgRequestVote MessageType = iota
	MsgAppendEntries
)

type Transport interface {
	Send(msg Message) error
	RegisterHandler(nodeID string, handler func(Message))
}
