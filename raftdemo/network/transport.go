package network

type Transport interface {
	Send(msg Message) error
	RegisterHandler(nodeID string, handler func(Message))
	Broadcast(msg Message) error
	Close() error
}
