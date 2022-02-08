package contracts

type Pipe func(passable interface{}) interface{}

type Pipeline interface {
	Send(passable interface{}) Pipeline
	Through(pipes ...interface{}) Pipeline
	Then(destination interface{}) interface{}
}
