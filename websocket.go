package contracts

type WebSocket interface {
	// Add 添加一个连接，返回 fd
	// add a connection, return fd.
	Add(connect WebSocketConnection)

	// GetFd 获取新的 fd
	// get new fd.
	GetFd() uint64

	// Close 关闭指定连接
	// close the specified connection.
	Close(fd uint64) error

	// Send 发送消息给指定 fd 的连接
	// send a message to the connection with the specified fd.
	Send(fd uint64, message any) error
}

type WebSocketConnection interface {
	WebSocketSender

	// Fd 获取 fd
	// get fd.
	Fd() uint64

	// Close 关闭该连接
	// close the connection.
	Close() error
}

type WebSocketSender interface {
	// Send 发送消息给该连接
	// send a message to this connection
	Send(message any) error

	// SendBytes 发送消息给该连接
	// send a message to this connection.
	SendBytes(bytes []byte) error

	// SendBinary 发送二进制消息给该连接
	// send binary message to this connection.
	SendBinary(bytes []byte) error
}

type WebSocketFrame interface {
	WebSocketSender

	// Connection 获取当前连接
	// get current connection.
	Connection() WebSocketConnection

	// Raw 获取原始消息
	// get the original message.
	Raw() []byte

	// RawString  获取字符串消息
	// get string message.
	RawString() string

	// Parse 解析 json 参数
	// Parse json parameters.
	Parse(v any) error
}

type WebSocketController interface {
	// OnConnect 可以在连接时处理一些鉴权之类的操作
	// Can handle some authentication and other operations when connecting.
	OnConnect(request HttpRequest, fd uint64) error

	// OnMessage 当有新的消息来时执行的操作
	// What to do when a new message arrives.
	OnMessage(frame WebSocketFrame)

	// OnClose 处理关闭事件
	// Handling close events.
	OnClose(fd uint64)
}
