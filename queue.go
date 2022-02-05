package contracts

type QueueFactory interface {
	Connection(name ...string) Queue
}

type QueueDriver func(config Fields) Queue

type Queue interface {
	// Size Get the size of the queue.

	Size() int64
	// Push a new job onto the queue.
	Push(job Job, data interface{}, queue ...string)

	// PushOn Push a new job onto the queue.
	PushOn(queue string, job Job, data interface{})

	// PushRaw Push a raw payload onto the queue.
	PushRaw(payload, queue string, options ...Fields)

	// Later Push a new job onto the queue after a delay.
	Later(delay interface{}, job Job, data interface{}, queue ...string)

	// LaterOn Push a new job onto the queue after a delay.
	LaterOn(queue string, delay interface{}, job Job, data interface{})

	// Bulk Push an array of jobs onto the queue.
	Bulk(job Job, data interface{}, queue ...string)

	// Pop the next job off of the queue.
	Pop(queue ...string) interface{}

	// GetConnectionName Get the connection name for the queue.
	GetConnectionName() string

	// SetConnectionName Set the connection name for the queue.
	SetConnectionName(queue string) Queue

	Listen() chan Job

	Stop()
}

type Job interface {

	// Uuid Get the UUID of the job.
	Uuid() string

	// GetJobId Get the job identifier.
	GetJobId() string

	// Payload Get the decoded body of the job.
	Payload() Fields

	// Fire the job.
	Fire()

	// Release
	/**
	 * Release the job back into the queue.
	 * Accepts a delay specified in seconds.
	 */
	Release(delay ...int)

	// IsReleased Determine if the job was released back into the queue.
	IsReleased() bool

	// Delete the job from the queue.
	Delete()

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

	// MaxTries Get the number of times to attempt a job.
	MaxTries() int

	// MaxExceptions Get the maximum number of exceptions allowed, regardless of attempts.
	MaxExceptions() int

	// Timeout Get the number of seconds the job can run.
	Timeout() int

	// RetryUntil Get the timestamp indicating when the job should timeout.
	RetryUntil() int

	// GetName Get the name of the queued job class.
	GetName() string

	// ResolveName
	/**
	 * Get the resolved name of the queued job class.
	 *
	 * Resolves the name of "wrapped" jobs such as class-based handlers.
	 */
	ResolveName() string

	// GetConnectionName Get the name of the connection the job belongs to.
	GetConnectionName() string

	// GetQueue Get the name of the queue the job belongs to.
	GetQueue() string

	// GetRawBody Get the raw body string for the job.
	GetRawBody() string
}

type QueueWorker interface {
	Work(queue Queue)
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
