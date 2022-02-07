package contracts

import "time"

type QueueFactory interface {
	Connection(name ...string) Queue
}

type Msg struct {
	Ack Ack
	Job Job
}

type Ack func()

type QueueDriver func(name string, config Fields, serializer JobSerializer) Queue

type Queue interface {
	// Push a new job onto the queue.
	Push(job Job, queue ...string)

	// PushOn Push a new job onto the queue.
	PushOn(queue string, job Job)

	// PushRaw Push a raw payload onto the queue.
	PushRaw(payload, queue string, options ...Fields) error

	// Later Push a new job onto the queue after a delay.
	Later(delay time.Time, job Job, queue ...string)

	// LaterOn Push a new job onto the queue after a delay.
	LaterOn(queue string, delay time.Time, job Job)

	// GetConnectionName Get the connection name for the queue.
	GetConnectionName() string

	// Release
	/**
	 * Release the job back into the queue.
	 * Accepts a delay specified in seconds.
	 */
	Release(job Job, delay ...int)

	// Delete the job from the queue.
	Delete(job Job)

	Listen(queue ...string) chan Msg

	Stop()
}

type Job interface {

	// Uuid Get the UUID of the job.
	Uuid() string

	// GetOptions Get the decoded body of the job.
	GetOptions() Fields

	// Handle the job.
	Handle()

	// IsReleased Determine if the job was released back into the queue.
	IsReleased() bool

	// IsDeleted Determine if the job has been deleted.
	IsDeleted() bool

	// IsDeletedOrReleased Determine if the job has been deleted or released.
	IsDeletedOrReleased() bool

	// Attempts Get the number of times the job has been attempted.
	Attempts() int

	// HasFailed Determine if the job has been marked as a failure.
	HasFailed() bool

	// MarkAsFailed Mark the job as "failed".
	MarkAsFailed()

	// Fail Delete the job, call the "failed" method, and raise the failed job event.
	Fail(err error)

	// GetMaxTries Get the max number of times to attempt a job.
	GetMaxTries() int

	// GetAttemptsNum Get the number of times to attempt a job.
	GetAttemptsNum() int

	// IncrementAttemptsNum increase the number of attempts
	IncrementAttemptsNum()

	// GetTimeout Get the number of seconds the job can run.
	GetTimeout() int

	// GetConnectionName Get the name of the connection the job belongs to.
	GetConnectionName() string

	// GetQueue Get the name of the queue the job belongs to.
	GetQueue() string

	SetQueue(queue string)
}

type QueueWorker interface {
	Work()
	Stop()
}

type JobSerializer interface {
	Serializer(job Job) string
	Unserialize(serialized string) (Job, error)
}

type ShouldQueue interface {
	ShouldQueue() bool
}

type ShouldBeUnique interface {
	ShouldBeUnique() bool
}

type ShouldBeEncrypted interface {
	ShouldBeEncrypted() bool
}
