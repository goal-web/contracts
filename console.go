package contracts

import "time"

type Console interface {
	// Call 按名称运行 控制台命令
	// run  console command by name.
	Call(command string, arguments CommandArguments) interface{}

	// Run 运行传入的控制台命令
	// run an incoming console command.
	Run(input ConsoleInput) interface{}

	// Schedule 注册任务调度
	// register Task Scheduler.
	Schedule(schedule Schedule)

	// GetSchedule 获取任务调度
	// get task schedule.
	GetSchedule() Schedule
}

type Command interface {
	// Handle 处理传入的控制台命令
	// handle an incoming console command.
	Handle() interface{}

	// InjectArguments 注入控制台命令参数
	// Inject console command parameters.
	InjectArguments(arguments CommandArguments) error

	// GetSignature 获取控制台命令签名
	// Get console command signature.
	GetSignature() string

	// GetName 获取控制台命令名称
	// Get console command name.
	GetName() string

	// GetDescription 获取控制台命令描述
	// Get console command description.
	GetDescription() string

	// GetHelp 获取控制台命令帮助信息
	// Get console command help
	GetHelp() string
}

type ConsoleInput interface {
	// GetCommand 获取控制台命令
	// get console command.
	GetCommand() string

	// GetArguments 获取控制台命令参数
	// Get console command parameters.
	GetArguments() CommandArguments
}

type CommandArguments interface {
	FieldsProvider
	Getter
	OptionalGetter

	// GetArg 获取控制台命令指定参数
	// Get the specified parameters of the console command.
	GetArg(index int) string

	// GetArgs 获取控制台命令参数
	// Get console command parameters.
	GetArgs() []string

	// SetOption 设置控制台命令选项
	// Set console command options.
	SetOption(key string, value interface{})

	// Exists 验证指定的键是否存在
	// Verify if the specified key exists.
	Exists(key string) bool

	// StringArrayOption 获取指定选项，返回[]string类型
	// Get the specified option, return []string type.
	StringArrayOption(key string, defaultValue []string) []string

	// IntArrayOption 获取指定选项，返回[]int类型
	// Get the specified option,  return []int type.
	IntArrayOption(key string, defaultValue []int) []int

	// Int64ArrayOption 获取指定选项，返回int64类型
	// Get the specified option, return int 64 type.
	Int64ArrayOption(key string, defaultValue []int64) []int64

	// FloatArrayOption 获取指定选项，返回[]float32类型
	// Get the specified option, return []float32 type.
	FloatArrayOption(key string, defaultValue []float32) []float32

	// Float64ArrayOption 获取指定选项，返回[]float64类型
	// Get the specified option, return []float64 type.
	Float64ArrayOption(key string, defaultValue []float64) []float64
}

type Schedule interface {

	// UseStore 使用指定的redis链接
	// Use the specified redis link.
	UseStore(store string)

	// Call 将新的回调事件添加到计划中
	// Add a new callback event to the schedule.
	Call(callback interface{}, args ...interface{}) CallbackEvent

	// Command 将新的命令事件添加到日程表
	// Add a new command event to the schedule.
	Command(command Command, args ...string) CommandEvent

	// Exec 将新的命令事件添加到计划中
	// Add a new command event to the schedule.
	Exec(command string, args ...string) CommandEvent

	// GetEvents 获取计划中的所有事件
	// Get all the events on the schedule.
	GetEvents() []ScheduleEvent
}

type ScheduleEvent interface {
	// Run 运行给定的事件
	// run the given event.
	Run(application Application)

	// WithoutOverlapping 不允许事件相互重叠
	// Do not allow the event to overlap each other.
	WithoutOverlapping(expiresAt int) ScheduleEvent

	// OnOneServer 允许每个 cron 表达式的事件仅在一台服务器上运行
	// Allow the event to only run on one server for each cron expression.
	OnOneServer() ScheduleEvent

	// MutexName 获取计划命令的互斥锁名称
	// Get the mutex name for the scheduled command.
	MutexName() string

	// SetMutexName 设置计划命令的互斥锁名称
	// Set the mutex name for the scheduled command.
	SetMutexName(mutexName string) ScheduleEvent



	// Skip 注册回调以进一步过滤计划, 返回true时跳过
	// Register a callback to further filter the schedule, skip when returning true.
	Skip(callback func() bool) ScheduleEvent

	// When 注册回调以进一步过滤计划, 返回true时不跳过
	// Register a callback to further filter the schedule, Do not skip when returning true.
	When(callback func() bool) ScheduleEvent



	// SpliceIntoPosition 将给定值拼接到表达式的给定位置
	// Splice the given value into the given position of the expression.
	SpliceIntoPosition(position int, value string) ScheduleEvent

	// Expression 获取事件的 cron 表达式
	// Get the cron expression for the event.
	Expression() string

	// Cron 表示事件频率的 cron 表达式
	// The cron expression representing the event's frequency.
	Cron(expression string) ScheduleEvent

	// Timezone 从字符串或对象设置实例的时区
	// Set the instance's timezone from a string or object.
	Timezone(timezone string) ScheduleEvent

	// Days 设置命令应该在一周中的哪几天运行
	// Set the days of the week the command should run on.
	Days(day string, days ...string) ScheduleEvent

	// Years 设置命令应该在哪几年运行
	// Set which years the command should run.
	Years(years ...string) ScheduleEvent

	// Yearly 安排活动每年运行
	// schedule the event to run yearly.
	Yearly() ScheduleEvent

	// YearlyOn 安排事件在给定的月份、日期和时间每年运行
	// schedule the event to run yearly on a given month, day, and time.
	YearlyOn(month time.Month, dayOfMonth int, time string) ScheduleEvent

	// Quarterly 安排活动每季度运行一次
	// schedule the event to run quarterly.
	Quarterly() ScheduleEvent

	// LastDayOfMonth 安排活动在每月的最后一天运行
	// schedule the event to run on the last day of the month.
	LastDayOfMonth(time string) ScheduleEvent

	// TwiceMonthly 安排活动在给定时间每月运行两次
	// schedule the event to run twice monthly at a given time.
	TwiceMonthly(first, second int, time string) ScheduleEvent

	// Monthly 安排活动每月运行
	// schedule the event to run monthly.
	Monthly() ScheduleEvent

	// MonthlyOn 安排活动在给定的日期和时间每月运行
	// schedule the event to run monthly on a given day and time.
	MonthlyOn(dayOfMonth int, time string) ScheduleEvent

	// WeeklyOn 安排活动在给定的日期和时间每周运行
	// schedule the event to run weekly on a given day and time.
	WeeklyOn(dayOfWeek time.Weekday, time string) ScheduleEvent

	// Weekly 安排活动每周运行
	// schedule the event to run weekly.
	Weekly() ScheduleEvent

	// Sundays 安排活动仅在周日运行
	// schedule the event to run only on sundays.
	Sundays() ScheduleEvent

	// Saturdays 安排活动仅在星期六运行
	// schedule the event to run only on saturdays.
	Saturdays() ScheduleEvent

	// Fridays 安排活动仅在周五运行
	// schedule the event to run only on fridays.
	Fridays() ScheduleEvent

	// Thursdays 安排活动仅在星期四运行
	// schedule the event to run only on thursdays.
	Thursdays() ScheduleEvent

	// Wednesdays 安排活动仅在星期三运行
	// schedule the event to run only on wednesdays.
	Wednesdays() ScheduleEvent

	// Tuesdays 安排活动仅在星期二运行
	// schedule the event to run only on tuesdays.
	Tuesdays() ScheduleEvent

	// Mondays 安排活动仅在星期一运行
	// schedule the event to run only on mondays.
	Mondays() ScheduleEvent

	// Weekends 安排活动仅在周末运行
	// schedule the event to run only on weekends.
	Weekends() ScheduleEvent

	// Weekdays 安排活动仅在工作日运行
	// schedule the event to run only on weekdays.
	Weekdays() ScheduleEvent

	// TwiceDailyAt 安排事件在给定的偏移量每天运行两次
	// schedule the event to run twice daily at a given offset.
	TwiceDailyAt(first, second, offset int) ScheduleEvent

	// TwiceDaily 安排活动每天运行两次。
	// schedule the event to run twice daily.
	TwiceDaily(first, second int) ScheduleEvent

	// DailyAt 安排活动在每天的给定时间（10:00、19:30 等）运行
	// schedule the event to run daily at a given time (10:00, 19:30, etc).
	DailyAt(time string) ScheduleEvent

	// Daily 安排活动每天运行
	// schedule the event to run daily.
	Daily() ScheduleEvent

	// EverySixHours 安排活动每六个小时运行一次
	// schedule the event to run every six hours.
	EverySixHours() ScheduleEvent

	// EveryFourHours 安排活动每四个小时运行一次
	// schedule the event to run every four hours.
	EveryFourHours() ScheduleEvent

	// EveryThreeHours 安排活动每三个小时运行一次
	// schedule the event to run every three hours.
	EveryThreeHours() ScheduleEvent

	// EveryTwoHours 安排活动每两个小时运行一次
	// schedule the event to run every two hours.
	EveryTwoHours() ScheduleEvent

	// HourlyAt 安排事件以每小时给定的偏移量每小时运行一次
	// schedule the event to run hourly at a given offset in the hour.
	HourlyAt(offset ...int) ScheduleEvent

	// Hourly 安排活动每小时运行一次
	// schedule the event to run hourly.
	Hourly() ScheduleEvent

	// EveryThirtyMinutes 安排活动每三十分钟运行一次
	// schedule the event to run every thirty minutes.
	EveryThirtyMinutes() ScheduleEvent

	// EveryFifteenMinutes 安排活动每十五分钟运行一次
	// schedule the event to run every fifteen minutes.
	EveryFifteenMinutes() ScheduleEvent

	// EveryTenMinutes 安排活动每十分钟运行一次
	// schedule the event to run every ten minutes.
	EveryTenMinutes() ScheduleEvent

	// EveryFiveMinutes 安排活动每五分钟运行一次
	// schedule the event to run every five minutes.
	EveryFiveMinutes() ScheduleEvent

	// EveryFourMinutes 安排活动每四分钟运行一次
	// schedule the event to run every four minutes.
	EveryFourMinutes() ScheduleEvent

	// EveryThreeMinutes 安排活动每三分钟运行一次
	// schedule the event to run every three minutes.
	EveryThreeMinutes() ScheduleEvent

	// EveryTwoMinutes 安排活动每两分钟运行一次
	// schedule the event to run every two minutes.
	EveryTwoMinutes() ScheduleEvent

	// EveryMinute 安排活动每分钟运行一次
	// schedule the event to run every minute.
	EveryMinute() ScheduleEvent

	// UnlessBetween 安排事件不在开始时间和结束时间之间运行
	// schedule the event to not run between start and end time.
	UnlessBetween(startTime, endTime string) ScheduleEvent

	// Between 安排事件在开始时间和结束时间之间运行
	// schedule the event to run between start and end time.
	Between(startTime, endTime string) ScheduleEvent


	// EveryThirtySeconds 安排活动每30秒运行一次
	// schedule the event to run every 30 seconds.
	EveryThirtySeconds() ScheduleEvent

	// EveryFifteenSeconds 安排活动每15秒运行一次
	// schedule the event to run every 15 seconds.
	EveryFifteenSeconds() ScheduleEvent

	// EveryTenSeconds 安排活动每10秒运行一次
	// schedule the event to run every 10 seconds.
	EveryTenSeconds() ScheduleEvent

	// EveryFiveSeconds 安排活动每5秒运行一次
	// schedule the event to run every 5 seconds.
	EveryFiveSeconds() ScheduleEvent

	// EveryFourSeconds 安排活动每4秒运行一次
	// schedule the event to run every 4 seconds.
	EveryFourSeconds() ScheduleEvent

	// EveryThreeSeconds 安排活动每3秒运行一次
	// schedule the event to run every 3 seconds.
	EveryThreeSeconds() ScheduleEvent

	// EveryTwoSeconds 安排活动每2秒运行一次
	// schedule the event to run every 2 seconds.
	EveryTwoSeconds() ScheduleEvent

	// EverySecond 安排活动每秒运行一次
	// schedule the event to run every second.
	EverySecond() ScheduleEvent
}

type CallbackEvent interface {
	ScheduleEvent

	// Description 设置事件的人性化描述
	// Set the human-friendly description of the event.
	Description(description string) CallbackEvent
}
type CommandEvent interface {
	ScheduleEvent
}
