package network

type MessageType int

const (
	MsgRequestVote MessageType = iota
	MsgRequestVoteResponse
	MsgAppendEntries
	MsgAppendResponse
)

type Message struct {
	From string
	To   string
	Type string // "AppendEntries", "AppendResponse", etc.
	Data interface{}
}
