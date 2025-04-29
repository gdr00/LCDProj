package consensus

import "sync"

type State int

const (
	Follower State = iota
	Leader
	Candidate
)

type ConsensusModule struct {
	mu       sync.Mutex
	state    State `default:"Follower"`
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

func (cm *ConsensusModule) GetPeers() []int {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	return cm.peers
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
