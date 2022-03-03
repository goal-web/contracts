package contracts

type HasherProvider func(config Fields) Hasher

type HasherFactory interface {
	Hasher

	// Driver 获取驱动程序实例
	// Get a driver instance.
	Driver(driver string) Hasher

	// Extend 注册一个自定义驱动程序提供者
	// Register a custom driver creator Closure.
	Extend(driver string, hasherProvider HasherProvider)
}

type Hasher interface {
	// Info 获取有关给定哈希值的信息
	// Get information about the given hashed value.
	Info(hashedValue string) Fields

	// Make 散列给定的值
	// Hash the given value.
	Make(value string, options Fields) string

	// Check 根据散列检查给定的普通值
	// check the given plain value against a hash.
	Check(value, hashedValue string, options Fields) bool
}
