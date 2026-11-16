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

// LogLevel represents the severity of a log record.
// LogLevel 表示日志记录的严重级别。
type LogLevel int

const (
	// DebugLevel is for verbose diagnostic output.
	// DebugLevel 用于详细的诊断输出。
	DebugLevel LogLevel = iota
	// InfoLevel is for general operational messages.
	// InfoLevel 用于常规运行信息。
	InfoLevel
	// WarnLevel is for abnormal but recoverable conditions.
	// WarnLevel 用于异常但可恢复的情况。
	WarnLevel
	// ErrorLevel is for errors that should be addressed.
	// ErrorLevel 用于需要处理的错误。
	ErrorLevel
	// FatalLevel is for unrecoverable errors, typically followed by process exit.
	// FatalLevel 用于不可恢复的错误，通常随后进程退出。
	FatalLevel
)

// ChannelLogger extends Logger with level-based dynamic logging.
// ChannelLogger 在 Logger 基础上支持按级别动态记录。
type ChannelLogger interface {
	Logger
	// Log adds a log record at the given level.
	// Log 按给定级别添加日志记录。
	Log(level LogLevel, msg string)
}

// LogManager resolves named log channels and registers custom drivers.
// LogManager 按名称解析日志通道，并注册自定义驱动。
type LogManager interface {
	// Channel returns the logger for the given channel; no name means the default channel.
	// Channel 返回给定通道的 logger；不传名称时返回默认通道。
	Channel(name ...string) ChannelLogger
	// Driver registers a custom log driver under the given name.
	// Driver 以给定名称注册自定义日志驱动。
	Driver(name string, driver LogDriver)
}

// LogDriver builds a Logger from a channel's configuration.
// LogDriver 按通道配置构建 Logger。
type LogDriver func(channelName string, config Fields) (Logger, error)
