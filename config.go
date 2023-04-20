package contracts

type ConfigProvider func(env Env) any

type Config interface {
	OptionalGetter[any]
	Getter[any]
	FieldsProvider

	// Reload 根据给定的字段提供者加载配置
	// reload configuration based on given field provider.
	Reload()

	// Set 设置给定的配置值
	// set a given configuration value.
	Set(key string, value any)

	// Unset 销毁指定的配置值
	// Destroy the specified configuration value.
	Unset(key string)
}

type Env interface {
	Getter[any]
	OptionalGetter[any]

	// Load 加载配置
	// load configuration.
	Load() Fields
}
