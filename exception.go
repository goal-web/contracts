package contracts

type Exception interface {
	Error() string
	GetPrevious() Exception
}

type ExceptionHandler interface {
	// Handle 处理异常
	// Handle the exception, and return the specified result.
	Handle(exception Exception) any

	// ShouldReport 判断是否需要上报
	// Determine whether to report.
	ShouldReport(exception Exception) bool

	// Report 上报异常
	// report exception.
	Report(exception Exception)
}
