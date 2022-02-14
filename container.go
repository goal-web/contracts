package contracts

import "reflect"

// InstanceProvider 容器实例提供者
// container instance provider
type InstanceProvider func() interface{}

type Container interface {

	// Bind 向容器注册绑定
	// Register a binding with the container.
	Bind(string, interface{})

	// Instance 将现有实例注册为在容器中共享
	// Register an existing instance as shared in the container.
	Instance(string, interface{})

	// Singleton 在容器中注册一个共享绑定
	// Register a shared binding in the container.
	Singleton(string, interface{})

	// HasBound 判断是否绑定
	// Determine whether to bind.
	HasBound(string) bool

	// Alias 将类型别名为不同的名称
	// alias a type to a different name.
	Alias(string, string)

	// Flush 刷新所有绑定和已解析实例的容器。
	// flush the container of all bindings and resolved instances.
	Flush()

	// Get 从容器中获取给定的类型
	// get the given type from the container.
	Get(key string, args ...interface{}) interface{}

	// Call 调用给定的 fn class@method 并注入其依赖项。
	// call the given fn / class@method and inject its dependencies.
	Call(fn interface{}, args ...interface{}) []interface{}

	// StaticCall 调用给定的 magical func class@method 并注入其依赖项。
	// call the given magical func / class@method and inject its dependencies.
	StaticCall(fn MagicalFunc, args ...interface{}) []interface{}

	// DI 从容器中注入给定的类型
	// injects the given type from the container.
	DI(object interface{}, args ...interface{})
}

// Component 可注入的类
// injectable class.
type Component interface {
	// Construct 构造函数
	// construct
	Construct(container Container)
}

// MagicalFunc 可以通过容器调用的魔术方法
// Magic methods that can be called from the container.
type MagicalFunc interface {
	// NumOut 输出参数个数
	// number of output parameters.
	NumOut() int

	// NumIn 输入参数个数
	// number of input parameters.
	NumIn() int

	// Call 调用
	// transfer.
	Call(in []reflect.Value) []reflect.Value

	// Arguments 获取所有参数
	// get all parameters.
	Arguments() []reflect.Type

	// Returns 获取所有返回类型
	// get all return types.
	Returns() []reflect.Type
}
