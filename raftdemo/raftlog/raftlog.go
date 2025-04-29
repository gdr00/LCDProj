package raftlog

type Entry struct {
	Term    int
	Command []byte
}

type Log interface {
	Append(entry Entry)
	Get(index int) (Entry, bool)
	LastIndex() int
	EntriesFrom(index int) []Entry
}
