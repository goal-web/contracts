package contracts

import (
	"mime/multipart"
	"net/http"
	"net/url"
)

type HttpEngine interface {
	Start(address string) error
	Close() error
	Static(prefix, directory string)
	Request() HttpRequest
}

type HttpContext interface {
}

type HttpResponse interface {
	// Status 获取 http 响应状态码
	// Get http response status code.
	Status() int

	Headers() http.Header

	Bytes() []byte
}

type HttpRequest interface {
	Getter[any]
	OptionalGetter[any]
	FieldsProvider

	Parse(form any) error

	// IsTLS 如果 HTTP 连接是 TLS，则返回 true，否则返回 false
	// returns true if HTTP connection is TLS otherwise false.
	IsTLS() bool

	// Scheme 返回 HTTP 协议方案，`http` 或 `https`
	// returns the HTTP protocol scheme, `http` or `https`.
	Scheme() string

	// RealIP 根据 `X-Forwarded-For` 返回客户端的网络地址或 `X-Real-IP` 请求标头，可以使用 `Echo#IPExtractor` 配置行为
	// returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	// The behavior can be configured using `Echo#IPExtractor`.
	RealIP() string

	// Path 返回处理程序的注册路径
	// returns the registered path for the handler.
	Path() string

	// Param 按名称返回路径参数
	// returns path parameter by name.
	Param(name string) string

	// QueryParam 返回提供的名称的查询参数
	// returns the query param for the provided name.
	QueryParam(name string) string

	// QueryParams 将查询参数返回为 `url.Values`
	// returns the query parameters as `url.Values`.
	QueryParams() url.Values

	// QueryString 返回 URL 查询字符串
	// returns the URL query string.
	QueryString() string

	// FormValue 返回提供的名称的表单字段值
	// returns the form field value for the provided name.
	FormValue(name string) string

	// FormParams 将表单参数返回为 `url.Values`
	// returns the form parameters as `url.Values`.
	FormParams() (Fields, error)

	// FormFile 返回所提供名称的多部分表单文件
	// returns the multipart form file for the provided name.
	FormFile(name string) (*multipart.FileHeader, error)

	// MultipartForm 返回多部分表单
	// returns the multipart form.
	MultipartForm() (*multipart.Form, error)

	// Cookie 返回请求中提供的命名cookie
	// returns the named cookie provided in the request.
	Cookie(name string) (string, error)

	// Cookies 返回随请求发送的 HTTP cookie
	// returns the HTTP cookies sent with the request.
	Cookies() []*http.Cookie

	SetCookie(*http.Cookie)

	// Get 从上下文中检索数据
	// retrieves data from the context.
	Get(key string) any

	// Set 在上下文中保存数据
	// saves data in the context.
	Set(key string, val any)

	Only(keys ...string) Fields
	Except(keys ...string) Fields

	GetHeader(key string) string

	SetHeader(key, value string)
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
