package contracts

import "time"

type QueueFactory interface {
	// Connection 解析队列连接实例
	// Resolve a queue connection instance.
	Connection(name ...string) Queue

	// Extend 添加队列连接解析器
	// Add a queue connection resolver.
	Extend(name string, driver QueueDriver)
}

type Msg struct {
	Ack Ack
	Job Job
}

type Ack func()


// QueueDriver 通过给定的信息获取队列连接
// Get queue connection with given information.
type QueueDriver func(name string, config Fields, serializer JobSerializer) Queue

type Queue interface {
	// Push 一个新工作进入队列
	// a new job onto the queue.
	Push(job Job, queue ...string) error

	// PushOn 将新作业推送到队列中
	// push a new job onto the queue.
	PushOn(queue string, job Job) error

	// PushRaw 将原始有效负载推送到队列中
	// push a raw payload onto the queue.
	PushRaw(payload, queue string, options ...Fields) error

	// Later 延迟后将新作业推送到队列中
	// push a new job onto the queue after a delay.
	Later(delay time.Time, job Job, queue ...string) error

	// LaterOn 延迟后将新作业推送到队列中
	// push a new job onto the queue after a delay.
	LaterOn(queue string, delay time.Time, job Job) error

	// GetConnectionName 获取队列的连接名称
	// Get the connection name for the queue.
	GetConnectionName() string

	// Release 将作业释放回队列中。接受以秒为单位指定的延迟
	// release the job back into the queue.
	// Accepts a delay specified in seconds.
	Release(job Job, delay ...int) error

	// Listen 监听给定的队列
	// listen to the given queue.
	Listen(queue ...string) chan Msg

	// Stop 关闭队列
	// close queue.
	Stop()
}

type Job interface {

	// Uuid 获取作业的 UUID
	// Get the UUID of the job.
	Uuid() string

	// GetOptions 获取作业的解码主体
	// Get the decoded body of the job.
	GetOptions() Fields

	// Handle 执行工作
	// the job.
	Handle()

	// IsReleased 确定作业是否被释放回队列
	// Determine if the job was released back into the queue.
	IsReleased() bool

	// IsDeleted 确定作业是否已被删除
	// Determine if the job has been deleted.
	IsDeleted() bool

	// IsDeletedOrReleased 确定作业是否已被删除或释放
	// Determine if the job has been deleted or released.
	IsDeletedOrReleased() bool

	// Attempts 获取作业已尝试的次数
	// Get the number of times the job has been attempted.
	Attempts() int

	// HasFailed 确定作业是否已被标记为失败
	// Determine if the job has been marked as a failure.
	HasFailed() bool

	// MarkAsFailed 将作业标记为“失败”
	// Mark the job as "failed".
	MarkAsFailed()

	// Fail 删除作业，调用“失败”方法，并引发失败的作业事件
	// Delete the job, call the "failed" method, and raise the failed job event.
	Fail(err error)

	// GetMaxTries 获取尝试工作的最大次数
	// Get the max number of times to attempt a job.
	GetMaxTries() int

	// GetRetryInterval 获取重试间隔任务之间的时间间隔，以秒为单位
	// Get the interval between retry interval tasks, in seconds
	GetRetryInterval() int

	// GetAttemptsNum 获取重试间隔任务之间的时间间隔，以秒为单位
	// Get the number of times to attempt a job.
	GetAttemptsNum() int

	// IncrementAttemptsNum 增加尝试次数
	// increase the number of attempts.
	IncrementAttemptsNum()

	// GetTimeout 获取作业可以运行的秒数
	// Get the number of seconds the job can run.
	GetTimeout() int

	// GetConnectionName 获取作业所属的连接名称
	// Get the name of the connection the job belongs to.
	GetConnectionName() string

	// GetQueue 获取作业所属队列的名称
	// Get the name of the queue the job belongs to.
	GetQueue() string

	// SetQueue 设置作业所属队列的名称
	// Sets the name of the queue to which the job belongs.
	SetQueue(queue string)
}

type QueueWorker interface {
	// Work 执行工作
	// perform work.
	Work()

	// Stop 停止工作
	// stop working.
	Stop()
}

type JobSerializer interface {
	// Serializer 将 "Job实例" 序列化为字符串
	// Serialize "Job instance" to string
	Serializer(job Job) string

	// Unserialize 将序列化后的字符串化为 "Job实例"
	// Convert the serialized string to a "Job instance"
	Unserialize(serialized string) (Job, error)
}

type ShouldQueue interface {
	// ShouldQueue 判断是否排队
	// Determine whether to queue.
	ShouldQueue() bool
}

type ShouldBeUnique interface {
	// ShouldBeUnique 判断是否唯一
	// determine whether it is unique.
	ShouldBeUnique() bool
}

type ShouldBeEncrypted interface {
	// ShouldBeEncrypted 判断是否加密
	// Determine whether to encrypt.
	ShouldBeEncrypted() bool
}
