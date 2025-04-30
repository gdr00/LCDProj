package network

type Transport interface {
	Send(msg Message) error
	Recive() (Message, error)
	Broadcast(msg Message) error
	Close() error
}
