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
	state    State
	leaderId int
	peers    []int
}

func (cm *ConsensusModule) SetState(state State) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.state = state
}

func (cm *ConsensusModule) SetLeaderId(leaderId int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.leaderId = leaderId
}

func (cm *ConsensusModule) SetPeers(peers []int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.peers = peers
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
