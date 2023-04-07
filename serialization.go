package contracts

type Serialization interface {
	// Method 根据给定的名称获取序列化驱动程序
	// Get the serialized driver based on the given name
	Method(name string) Serializer

	// Extend 添加序列化驱动程序
	// Add serialization driver.
	Extend(name string, serializer Serializer)
}

type Serializer interface {
	// Serialize 序列化给定的数据
	// serialize the given data.
	Serialize(any) string

	// Unserialize 反序列化
	// deserialize.
	Unserialize(string, any) error
}

type ClassSerializer interface {
	// Register 注册解析类
	// register parsing class.
	Register(class Class[any])

	// Serialize 序列化给定的数据
	// serialize the given data.
	Serialize(any) string

	// Parse 解析序列化后的字符串
	// parse the serialized string.
	Parse(string) (any, error)
}
