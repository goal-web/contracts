package contracts

type Pipeline interface {
	Send(passable interface{}) Pipeline
	Through(pipes ...interface{}) Pipeline
	Then(destination interface{}) interface{}
}
