package network

type MessageType int

type Message struct {
	From    string
	To      string
	Type    MessageType
	Payload interface{}
}
