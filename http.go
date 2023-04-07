package contracts

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type HttpContext interface {
	// Request 返回`*http.Request`
	// returns `*http.Request`.
	Request() *http.Request

	// SetRequest 设置`*http.Request`
	// sets `*http.Request`.
	SetRequest(r *http.Request)

	// IsTLS 如果 HTTP 连接是 TLS，则返回 true，否则返回 false
	// returns true if HTTP connection is TLS otherwise false.
	IsTLS() bool

	// IsWebSocket 如果 HTTP 连接是 websocket，则返回 true，否则返回 false
	// returns true if HTTP connection is websocket otherwise false.
	IsWebSocket() bool

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

	// SetPath 设置处理程序的注册路径
	// sets the registered path for the handler.
	SetPath(p string)

	// Param 按名称返回路径参数
	// returns path parameter by name.
	Param(name string) string

	// ParamNames 返回路径参数名称
	// returns path parameter names.
	ParamNames() []string

	// SetParamNames 设置路径参数名称
	// sets path parameter names.
	SetParamNames(names ...string)

	// ParamValues 返回路径参数值
	// returns path parameter values.
	ParamValues() []string

	// SetParamValues 设置路径参数值
	// sets path parameter values.
	SetParamValues(values ...string)

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
	FormParams() (url.Values, error)

	// FormFile 返回所提供名称的多部分表单文件
	// returns the multipart form file for the provided name.
	FormFile(name string) (*multipart.FileHeader, error)

	// MultipartForm 返回多部分表单
	// returns the multipart form.
	MultipartForm() (*multipart.Form, error)

	// Cookie 返回请求中提供的命名cookie
	// returns the named cookie provided in the request.
	Cookie(name string) (*http.Cookie, error)

	// SetCookie 在 HTTP 响应中添加一个 `Set-Cookie` 标头
	// adds a `Set-Cookie` header in HTTP response.
	SetCookie(cookie *http.Cookie)

	// Cookies 返回随请求发送的 HTTP cookie
	// returns the HTTP cookies sent with the request.
	Cookies() []*http.Cookie

	// Get 从上下文中检索数据
	// retrieves data from the context.
	Get(key string) any

	// Set 在上下文中保存数据
	// saves data in the context.
	Set(key string, val any)

	// Bind 将请求正文绑定到提供的类型“i”。默认绑定器基于 Content-Type 标头执行此操作
	// binds the request body into provided type `i`. The default
	// binder does it based on Content-Type header.
	Bind(i any) error

	// Validate 验证提供的`i`。它通常在 `HttpContextBind()` 之后调用，验证器必须使用 `Echo#Validator` 注册
	// validates provided `i`. It is usually called after `HttpContext#Bind()`.
	// Validator must be registered using `Echo#Validator`.
	Validate(i any) error

	// Render 呈现带有数据的模板并发送带有状态码的 text/html 响应。渲染器必须使用 `Echo.Renderer` 注册
	// renders a template with data and sends a text/html response with status
	// code. Renderer must be registered using `Echo.Renderer`.
	Render(code int, name string, data any) error

	// HTML 发送带有状态码的 HTTP 响应
	// sends an HTTP response with status code.
	HTML(code int, html string) error

	// HTMLBlob 发送带有状态码的 HTTP blob 响应
	// sends an HTTP blob response with status code.
	HTMLBlob(code int, b []byte) error

	// String 发送带有状态码的字符串响应
	// sends a string response with status code.
	String(code int, s string) error

	// JSON 发送带有状态码的 json 响应
	// sends a json response with status code.
	JSON(code int, i any) error

	// JSONPretty 发送带有状态码的漂亮打印 json
	// sends a pretty-print json with status code.
	JSONPretty(code int, i any, indent string) error

	// JSONBlob 发送带有状态码的 json blob 响应
	// sends a json blob response with status code.
	JSONBlob(code int, b []byte) error

	// JSONP 发送带有状态码的 jsonp 响应。它使用 `callback` 来构造 jsonp 有效负载
	// sends a jsonp response with status code. It uses `callback` to construct
	// the jsonp payload.
	JSONP(code int, callback string, i any) error

	// JSONPBlob 发送带有状态码的 jsonp blob 响应。它使用 `callback` 构造 jsonp 有效负载
	// sends a jsonp blob response with status code. It uses `callback`
	// to construct the jsonp payload.
	JSONPBlob(code int, callback string, b []byte) error

	// XML 发送带有状态码的 xml 响应
	// sends an xml response with status code.
	XML(code int, i any) error

	// XMLPretty 发送带有状态码的漂亮打印 xml
	// sends a pretty-print xml with status code.
	XMLPretty(code int, i any, indent string) error

	// XMLBlob 发送带有状态码的 xml blob 响应
	// sends an xml blob response with status code.
	XMLBlob(code int, b []byte) error

	// Blob 发送带有状态代码和内容类型的 blob 响应
	// sends a blob response with status code and content type.
	Blob(code int, contentType string, b []byte) error

	// Stream 发送带有状态码和内容类型的流式响应
	// sends a streaming response with status code and content type.
	Stream(code int, contentType string, r io.Reader) error

	// File 发送包含文件内容的响应
	// sends a response with the content of the file.
	File(file string) error

	// Attachment 作为附件发送响应，提示客户端保存文件
	// sends a response as attachment, prompting client to save the file.
	Attachment(file string, name string) error

	// Inline 以内联形式发送响应，在浏览器中打开文件
	// sends a response as inline, opening the file in the browser.
	Inline(file string, name string) error

	// NoContent 发送没有正文和状态码的响应
	// sends a response with no body and a status code.
	NoContent(code int) error

	// Redirect 将请求重定向到提供的带有状态码的 URL
	// redirects the request to a provided URL with status code.
	Redirect(code int, url string) error

	// Error 调用已注册的 HTTP 错误处理程序。一般由中间件使用
	// invokes the registered HTTP error handler. Generally used by middleware.
	Error(err error)

	// Reset 请求完成后重置上下文。它必须与 `EchoAcquireContext()` 和 `EchoReleaseContext()` 一起调用。参见`EchoServeHTTP()`
	// resets the context after request completes. It must be called along
	// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
	// See `Echo#ServeHTTP()`
	Reset(r *http.Request, w http.ResponseWriter)
}

type HttpResponse interface {
	// Status 获取 http 响应状态码
	// Get http response status code.
	Status() int

	// Response 获取 http 响应
	// get http response.
	Response(ctx HttpContext) error
}

type HttpRequest interface {
	Getter[any]
	OptionalGetter[any]
	FieldsProvider

	// Only 只获取指定 key 的数据
	// Get only the data of the specified key.
	Only(keys ...string) Fields

	// OnlyExists 只获取指定 key ，不存在或者 nil 则忽略
	// Get only the specified key , ignore if it does not exist or nil.
	OnlyExists(keys ...string) Fields

	// Request 返回`*http.Request`
	// returns `*http.Request`.
	Request() *http.Request

	// SetRequest 设置 `*http.Request`
	// sets `*http.Request`.
	SetRequest(r *http.Request)

	// IsTLS 如果 HTTP 连接是 TLS，则返回 true，否则返回 false
	// returns true if HTTP connection is TLS otherwise false.
	IsTLS() bool

	// IsWebSocket 如果 HTTP 连接是 websocket，则返回 true，否则返回 false
	// returns true if HTTP connection is websocket otherwise false.
	IsWebSocket() bool

	// Scheme 返回 HTTP 协议方案，`http` 或 `https`
	// returns the HTTP protocol scheme, `http` or `https`.
	Scheme() string

	// RealIP 根据 `X-Forwarded-For` 或 `X-Real-IP` 请求头返回客户端的网络地址。可以使用 `Echo#IPExtractor` 配置行为
	// returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	// The behavior can be configured using `Echo#IPExtractor`.
	RealIP() string

	// Path 返回处理程序的注册路径
	// returns the registered path for the handler.
	Path() string

	// SetPath 设置处理程序的注册路径
	// sets the registered path for the handler.
	SetPath(p string)

	// Param 按名称返回路径参数
	// returns path parameter by name.
	Param(name string) string

	// ParamNames 返回路径参数名称
	// returns path parameter names.
	ParamNames() []string

	// SetParamNames 设置路径参数名称
	// sets path parameter names.
	SetParamNames(names ...string)

	// ParamValues 返回路径参数值
	// returns path parameter values.
	ParamValues() []string

	// SetParamValues 设置路径参数值
	// sets path parameter values.
	SetParamValues(values ...string)

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
	FormParams() (url.Values, error)

	// FormFile 返回所提供名称的多部分表单文件
	// returns the multipart form file for the provided name.
	FormFile(name string) (*multipart.FileHeader, error)

	// MultipartForm 返回多部分表单
	// returns the multipart form.
	MultipartForm() (*multipart.Form, error)

	// Cookie 返回请求中提供的命名cookie
	// returns the named cookie provided in the request.
	Cookie(name string) (*http.Cookie, error)

	// SetCookie 在 HTTP 响应中添加一个 `Set-Cookie` 标头
	// adds a `Set-Cookie` header in HTTP response.
	SetCookie(cookie *http.Cookie)

	// Cookies 返回随请求发送的 HTTP cookie
	// returns the HTTP cookies sent with the request.
	Cookies() []*http.Cookie

	// Get 从上下文中检索数据
	// retrieves data from the context.
	Get(key string) any

	// Set 在上下文中保存数据
	// saves data in the context.
	Set(key string, val any)

	// Bind 将请求正文绑定到提供的类型“i”。默认绑定器基于 Content-Type 标头执行此操作
	// binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i any) error

	// Validate 验证提供的`i`。它通常在 `HttpContext#Bind()` 之后调用。验证器必须使用 `Echo#Validator` 注册
	// validates provided `i`. It is usually called after `HttpContext#Bind()`.
	// Validator must be registered using `Echo#Validator`.
	Validate(i any) error

	// Render 呈现带有数据的模板并发送带有状态码的 text/html 响应。渲染器必须使用 `Echo.Renderer` 注册
	// renders a template with data and sends a text/html response with status
	// code. Renderer must be registered using `Echo.Renderer`.
	Render(code int, name string, data any) error

	// HTML 发送带有状态码的 HTTP 响应
	// sends an HTTP response with status code.
	HTML(code int, html string) error

	// HTMLBlob 发送带有状态码的 HTTP blob 响应
	// sends an HTTP blob response with status code.
	HTMLBlob(code int, b []byte) error

	// String 发送带有状态码的字符串响应
	// sends a string response with status code.
	String(code int, s string) error

	// JSON 发送带有状态码的 json 响应
	// sends a json response with status code.
	JSON(code int, i any) error

	// JSONPretty 发送带有状态码的漂亮打印 json
	// sends a pretty-print json with status code.
	JSONPretty(code int, i any, indent string) error

	// JSONBlob 发送带有状态码的 json blob 响应
	// sends a json blob response with status code.
	JSONBlob(code int, b []byte) error

	// JSONP 发送带有状态码的 jsonp 响应。它使用“回调”来构造 jsonp 有效负载
	// sends a jsonp response with status code. It uses `callback` to construct
	// the jsonp payload.
	JSONP(code int, callback string, i any) error

	// JSONPBlob 发送带有状态码的 jsonp blob 响应。它使用 `callback` 来构造 jsonp 有效负载
	// sends a jsonp blob response with status code. It uses `callback`
	// to construct the jsonp payload.
	JSONPBlob(code int, callback string, b []byte) error

	// XML 发送带有状态码的 xml 响应
	// sends an xml response with status code.
	XML(code int, i any) error

	// XMLPretty 发送带有状态码的漂亮打印 xml
	// sends a pretty-print xml with status code.
	XMLPretty(code int, i any, indent string) error

	// XMLBlob 发送带有状态码的 xml blob 响应
	// sends an xml blob response with status code.
	XMLBlob(code int, b []byte) error

	// Blob 发送带有状态代码和内容类型的 blob 响应
	// sends a blob response with status code and content type.
	Blob(code int, contentType string, b []byte) error

	// Stream 发送带有状态码和内容类型的流式响应
	// sends a streaming response with status code and content type.
	Stream(code int, contentType string, r io.Reader) error

	// File 发送包含文件内容的响应
	// sends a response with the content of the file.
	File(file string) error

	// Attachment 以附件形式发送响应，提示客户端保存文件
	// sends a response as attachment, prompting client to save the file.
	Attachment(file string, name string) error

	// Inline 以内联形式发送响应，在浏览器中打开文件
	// sends a response as inline, opening the file in the browser.
	Inline(file string, name string) error

	// NoContent 发送没有正文和状态码的响应
	// sends a response with no body and a status code.
	NoContent(code int) error

	// Redirect 将请求重定向到提供的带有状态码的 URL
	// redirects the request to a provided URL with status code.
	Redirect(code int, url string) error

	// Error 调用已注册的 HTTP 错误处理程序。一般由中间件使用
	// invokes the registered HTTP error handler. Generally used by middleware.
	Error(err error)

	// Reset 请求完成后重置上下文。它必须与 `Echo#AcquireContext()` 和 `Echo#ReleaseContext()` 一起调用。参见`Echo#ServeHTTP()`
	// resets the context after request completes. It must be called along
	// with `Echo#AcquireContext()` and `Echo#ReleaseContext()`.
	// See `Echo#ServeHTTP()`
	Reset(r *http.Request, w http.ResponseWriter)
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
