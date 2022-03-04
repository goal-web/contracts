package contracts

type Exception interface {
	error
	FieldsProvider
}

type ExceptionHandler interface {
	// Handle 处理异常
	// Handle the exception, and return the specified result.
	Handle(exception Exception) interface{}

	// ShouldReport 判断是否需要上报
	// Determine whether to report.
	ShouldReport(exception Exception) bool

	// Report 上报异常
	// report exception.
	Report(exception Exception)
}
