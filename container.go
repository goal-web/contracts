package contracts

// InstanceProvider 实例提供者
type InstanceProvider func() interface{}

// Container 容器
type Container interface {
	// Provide func() T
	Provide(interface{})
	// ProvideSingleton func() T
	ProvideSingleton(interface{})
	// Bind 绑定实例提供者
	Bind(string, InstanceProvider)
	// Instance 绑定实例
	Instance(string, interface{})
	// Singleton 绑定实例提供者，单例模式
	Singleton(string, InstanceProvider)
	// HasBound 判断是否绑定某 key
	HasBound(string) bool
	// Alias 设置别名
	Alias(string, string)
	// Flush 清除所有绑定
	Flush()
	// Get 获取指定实例
	Get(string) interface{}
	// Call 使用容器调用函数，支持依赖注入
	Call(interface{}, ...interface{}) []interface{}
	// DI 为结构体执行依赖注入
	DI(object interface{}, args ...interface{})
}

// Component 支持依赖注入的接口
type Component interface {
	ShouldInject()
}
