package consensus

import "sync"

type State int

const (
	Leader State = iota
	Follower
	Candidate
)

type ConsensusModule struct {
	mu       sync.Mutex
	id       int
	state    State
	leaderId int
	peers    []int
}

func (cm *ConsensusModule) IsLeader() bool {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	return cm.state == Leader
}

func (cm *ConsensusModule) GetLeader() int {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	return cm.leaderId
}
