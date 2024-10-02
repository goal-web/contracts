package contracts

type Session interface {
	// GetName 获取会话的名称
	// get the name of the session.
	GetName() string

	// SetName 设置会话的名称
	// Set the name of the session.
	SetName(name string)

	// GetId 获取当前会话 ID
	// get the current session ID.
	GetAuthenticatableKey() string

	// SetId 设置会话 ID
	// Set the session ID.
	SetId(id string)

	// Start 启动会话，从处理程序中读取数据
	// start the session, reading the data from a handler.
	Start() bool

	// Save 将会话数据保存到存储中
	// save the session data to storage.
	Save()

	// All 获取所有会话数据
	// get all of the session data.
	All() map[string]string

	// Exists 检查密钥是否存在
	// Checks if a key exists.
	Exists(key string) bool

	// Has 检查密钥是否存在且不为空
	// Checks if a key is present and not null.
	Has(key string) bool

	// Get 从会话中获取一个项目
	// get an item from the session.
	Get(key, defaultValue string) string

	// Pull 获取给定键的值，然后忘记它
	// get the value of a given key and then forget it.
	Pull(key, defaultValue string) string

	// Put 在会话中放置一个键值对或键值对数组
	// put a key / value pair or array of key / value pairs in the session.
	Put(key, value string)

	// Token 获取 CSRF 令牌值
	// get the CSRF token value.
	Token() string

	// RegenerateToken 重新生成 CSRF 令牌值
	// regenerate the CSRF token value.
	RegenerateToken()

	// Remove 从会话中删除一个项目，返回它的值
	// remove an item from the session, returning its value.
	Remove(key string) string

	// Forget 从会话中删除一项或多项
	// remove one or many items from the session.
	Forget(keys ...string)

	// Flush 从会话中删除所有项目
	// remove all of the items from the session.
	Flush()

	// Invalidate 刷新会话数据并重新生成 ID
	// flush the session data and regenerate the ID.
	Invalidate() bool

	// Regenerate 生成新的会话标识符
	// Generate a new session identifier.
	Regenerate(destroy bool) bool

	// Migrate 为会话生成一个新的会话 ID
	// Generate a new session ID for the session.
	Migrate(destroy bool) bool

	// IsStarted 确定会话是否已启动
	// Determine if the session has been started.
	IsStarted() bool

	// PreviousUrl 从会话中获取上一个 URL
	// get the previous URL from the session.
	PreviousUrl() string

	// SetPreviousUrl 在会话中设置“上一个” URL
	// Set the "previous" URL in the session.
	SetPreviousUrl(url string)
}

type SessionStore interface {
	// LoadSession 从处理程序加载会话数据
	// Load the session data from the handler.
	LoadSession(id string) map[string]string

	// Save 将会话数据保存到存储中
	// save the session data to storage.
	Save(id string, sessions map[string]string)
}
