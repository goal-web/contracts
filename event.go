package contracts

type Event interface {
	// Event 获取事件名称
	// get event name.
	Event() string
}

type SyncEvent interface {
	Event
	// Sync 判断是否同步事件
	// Determine whether to synchronize events.
	Sync() bool
}

type EventListener interface {
	// Handle 事件触发时处理事件，所有监听器都会
	// handle the event when the event is triggered, all listeners will
	Handle(event Event)
}

type EventDispatcher interface {
	// Register 向调度程序注册事件侦听器
	// Register an event listener with the dispatcher.
	Register(name string, listener EventListener)

	// Dispatch 为所有侦听器提供要处理的事件
	// Provide all listeners with an event to process.
	Dispatch(event Event)
}
