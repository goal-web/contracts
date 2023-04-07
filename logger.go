package contracts

type Logger interface {
	// WithFields 添加数据
	// adding data
	WithFields(fields Fields) Logger

	// WithField 通过给定的key value 添加数据
	// Add data by given key value.
	WithField(key string, value any) Logger

	// WithError 添加错误
	// add error.
	WithError(err error) Logger

	// WithException 将异常管理委托给自定义异常处理程序
	// Delegate exception management to a custom exception handler.
	WithException(exception Exception) Logger

	// Info 在 INFO 级别添加日志记录
	// Adds a log record at the INFO level.
	Info(msg string)

	// Warn 在 WARNING 级别添加日志记录
	// Adds a log record at the WARNING level.
	Warn(msg string)

	// Debug 在 DEBUG 级别添加日志记录
	// Adds a log record at the DEBUG level.
	Debug(msg string)

	// Error 在 ERROR 级别添加日志记录
	// Adds a log record at the ERROR level.
	Error(msg string)

	// Fatal 在 FATAL 级别添加日志记录
	// Adds a log record at the FATAL level.
	Fatal(msg string)
}
