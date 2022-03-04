package contracts

// Pipe 通过可用管道之一发送对象
// Send an object through one of the available pipelines.
type Pipe func(passable interface{}) interface{}

type Pipeline interface {
	// Send 设置通过管道发送的对象
	// Set the object being sent through the pipeline.
	Send(passable interface{}) Pipeline

	// Through 设置管道数组
	// Set the array of pipes.
	Through(pipes ...interface{}) Pipeline

	// Then 使用最终目标回调运行管道
	// Run the pipeline with a final destination callback.
	Then(destination interface{}) interface{}
}
