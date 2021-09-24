package contracts

// Exception 异常
type Exception interface {
	error
	Fields() Fields
}

// ExceptionHandler 异常处理器
type ExceptionHandler interface {
	// Handle 处理异常
	Handle(exception Exception)

	// ShouldReport 判断是否需要上报
	ShouldReport(exception Exception) bool

	// Report 上报异常
	Report(exception Exception)
}
