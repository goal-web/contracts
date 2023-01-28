package contracts

// Application 应用程序接口
// application interface
type Application interface {
	Container

	// GetExceptionHandler 获取异常处理器
	GetExceptionHandler() ExceptionHandler

	// GetConfig 获取配置类
	GetConfig() Config

	// IsProduction 判断是否为生产环境.
	// Determine if it is a production environment.
	IsProduction() bool

	// Debug 确定是否启用调试。
	// 如果启用调试，日志会打印一些调试信息。
	// Determine whether to enable debugging.
	// If debugging is enabled, the log will print some debugging information.
	Debug() bool

	// Environment 获取当前运行环境
	// Get the current operating environment.
	Environment() string

	// RegisterServices 注册应用服务.
	// Register the application service.
	RegisterServices(provider ...ServiceProvider)

	// Start 开启应用程序
	// application start.
	Start() map[string]error

	// Stop 关闭应用程序
	// application stop.
	Stop()
}

// ServiceProvider 服务提供者接口
// Service Provider Interface.
type ServiceProvider interface {
	// Register 注册服务
	// register the services.
	Register(application Application)

	// Start 启动服务
	// start service.
	Start() error

	// Stop 关闭服务
	// stop service.
	Stop()
}
