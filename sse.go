package contracts

type SseFactory interface {
	Sse(key string) Sse
	Register(key string, sse Sse)
}

type Sse interface {
	// Add 添加一个连接
	// add a connection.
	Add(connect SseConnection)

	// GetFd 获取新的 fd
	// get new fd.
	GetFd() uint64

	// Close 关闭 fd
	// close fd.
	Close(fd uint64) error

	// Send 发送消息给指定 fd 的连接
	// send a message to the connection with the specified fd.
	Send(fd uint64, message any) error

	// Count 返回
	Count() uint64

	// Broadcast 给所有连接到指定 sse 接口的客户端发送消息
	// send a message to everyone who connected this sse
	Broadcast(message any) error
}

type SseConnection interface {
	// Fd 获取连接标识
	// Get connection ID.
	Fd() uint64

	// Send 发送消息给该连接
	// send a message to this connection.
	Send(message any) error

	// Close 关闭连接标识
	// close connection ID.
	Close() error
}

type SseController interface {
	// OnConnect 可以在连接时处理一些鉴权之类的操作
	// Can handle some authentication and other operations when connecting.
	OnConnect(request HttpRequest, fd uint64) error

	// OnClose 处理关闭事件
	// Handling close events.
	OnClose(fd uint64)
}
