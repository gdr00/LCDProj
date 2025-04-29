package network

type MessageType int

const (
	MsgRequestVote MessageType = iota
	MsgAppendEntries
)

type Message struct {
	From    string
	To      string
	Type    MessageType
	Payload interface{}
}

type Transport interface {
	Send(msg Message) error
	RegisterHandler(nodeID string, handler func(Message))
}
