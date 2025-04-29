package statemachine

type StateMachine interface {
	Apply(command []byte) error
}
