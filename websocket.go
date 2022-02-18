package contracts

type WebSocket interface {
	// Add 添加一个连接，返回 fd
	Add(connect WebSocketConnection)

	// GetFd 获取新的 fd
	GetFd() uint64

	// Close 关闭指定连接
	Close(fd uint64) error

	// Send 发送消息给指定 fd 的连接
	Send(fd uint64, message interface{}) error
}

type WebSocketConnection interface {
	WebSocketSender

	// Fd 获取 fd
	Fd() uint64
	// Close 关闭该连接
	Close() error
}

type WebSocketSender interface {
	// Send 发送消息给该连接
	Send(message interface{}) error

	// SendBytes 发送消息给该连接
	SendBytes(bytes []byte) error

	// SendBinary 发送二进制消息给该连接
	SendBinary(bytes []byte) error
}

type WebSocketFrame interface {
	WebSocketSender

	// Connection 获取当前连接
	Connection() WebSocketConnection

	// Raw 获取原始消息
	Raw() []byte
	// RawString  获取字符串消息
	RawString() string

	// Parse 解析 json 参数
	Parse(v interface{}) error
}

type WebSocketController interface {
	// OnConnect 可以在连接时处理一些鉴权之类的操作
	OnConnect(request HttpRequest, fd uint64) error

	// OnMessage 当有新的消息来时执行的操作
	OnMessage(frame WebSocketFrame)

	// OnClose 处理关闭事件
	OnClose(fd uint64)
}
