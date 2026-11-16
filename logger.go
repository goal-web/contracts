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

// LogLevel 表示日志记录的严重级别。
// LogLevel represents the severity of a log record.
type LogLevel int

// 日志级别常量。零值为 DebugLevel（最详细的级别），因此配置缺失时会隐式落为 Debug 输出。
// Log level constants. The zero value is DebugLevel (the most verbose level), so a missing
// configuration implicitly falls back to Debug output.
const (
	// DebugLevel 用于详细的诊断输出。
	// DebugLevel is for verbose diagnostic output.
	DebugLevel LogLevel = iota
	// InfoLevel 用于常规运行信息。
	// InfoLevel is for general operational messages.
	InfoLevel
	// WarnLevel 用于异常但可恢复的情况。
	// WarnLevel is for abnormal but recoverable conditions.
	WarnLevel
	// ErrorLevel 用于需要处理的错误。
	// ErrorLevel is for errors that should be addressed.
	ErrorLevel
	// FatalLevel 用于不可恢复的错误，通常随后进程退出。
	// FatalLevel is for unrecoverable errors, typically followed by process exit.
	FatalLevel
)

// ChannelLogger 在 Logger 基础上支持按级别动态记录。
// ChannelLogger extends Logger with level-based dynamic logging.
type ChannelLogger interface {
	Logger
	// Log 按给定级别添加日志记录。
	// Log adds a log record at the given level.
	Log(level LogLevel, msg string)
}

// LogManager 按名称解析日志通道，并注册自定义驱动。
// LogManager resolves named log channels and registers custom drivers.
type LogManager interface {
	// Channel 返回给定通道的 logger；不传名称时返回默认通道。
	// Channel returns the logger for the given channel; no name means the default channel.
	Channel(name ...string) ChannelLogger
	// Extend 以给定名称注册自定义日志驱动。
	// Extend registers a custom log driver under the given name.
	Extend(name string, driver LogDriver)
}

// LogDriver 按通道配置构建 Logger。
// 框架会将返回的 Logger 包装为 ChannelLogger；原生支持按级别记录的驱动可直接实现 ChannelLogger。
// LogDriver builds a Logger from a channel's configuration.
// The framework wraps the returned Logger as a ChannelLogger; drivers with native
// level-based logging may implement ChannelLogger directly.
type LogDriver func(channelName string, config Fields) (Logger, error)
