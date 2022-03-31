package contracts

// GuardDriver 守卫驱动程序
// guard driver
type GuardDriver func(name string, config Fields, ctx Context, provider UserProvider) Guard

// UserProviderDriver 用户提供者驱动程序
// User Provider Driver
type UserProviderDriver func(config Fields) UserProvider

type Auth interface {

	// ExtendUserProvider 扩展用户提供者
	// Extended User Provider.
	ExtendUserProvider(name string, provider UserProviderDriver)

	// ExtendGuard 扩展守卫
	// Extended guard.
	ExtendGuard(name string, guard GuardDriver)

	// Guard 按名称获取守卫实例
	// Get a guard instance by name.
	Guard(name string, ctx Context) Guard

	// UserProvider 按名称获取用户提供者实例
	// Get a user provider instance by name.
	UserProvider(name string) UserProvider
}

type Authenticatable interface {
	// GetId 获取当前已认证用户的 ID
	// Get the ID for the currently authenticated user.
	GetId() string
}

type Guard interface {
	// Once 设置当前用户
	// Set the current user.
	Once(user Authenticatable)

	// User 获取当前认证的用户
	// Get the currently authenticated user.
	User() Authenticatable

	// GetId 获取当前已认证用户的 ID
	// Get the ID for the currently authenticated user.
	GetId() string

	// Check 判断当前用户是否经过身份验证
	// Determine if the current user is authenticated.
	Check() bool

	// Guest 判断当前用户是否为访客
	// Determine if the current user is a guest.
	Guest() bool

	// Login 将用户登录到应用程序
	// Log a user into the application.
	Login(user Authenticatable) interface{}

	// Logout 用户登出
	Logout() error
}

type UserProvider interface {
	// RetrieveById 通过用户的唯一标识符检索用户
	// Retrieve a user by their unique identifier.
	RetrieveById(identifier string) Authenticatable
}

type Authorizable interface {
	// Can 确定实体是否具有给定的能力
	// Determine if the entity has a given ability.
	Can(ability string, arguments ...interface{}) bool
}

// GateChecker 权限检查器
// permission checker.
type GateChecker func(user Authorizable, data ...interface{}) bool

// GateHook 权限钩子
// permission hook.
type GateHook func(user Authorizable, ability string, data ...interface{}) bool

// Policy 权限策略, 一组检查器
// Permission policy, a set of checkers.
type Policy map[string]GateChecker

type Gate interface {

	// Allows 确定是否应该为当前用户授予给定的能力
	// determined if the given ability should be granted for the current user.
	Allows(ability string, arguments ...interface{}) bool

	// Denies 确定是否应该为当前用户拒绝给定的能力
	// Determine if the given ability should be denied for the current user.
	Denies(ability string, arguments ...interface{}) bool

	// Check 确定是否应为当前用户授予所有给定的能力
	// Determine if all the given abilities should be granted for the current user.
	Check(abilities []string, arguments ...interface{}) bool

	// Any 确定是否应为当前用户授予任何一种给定能力
	// Determine if any one of the given abilities should be granted for the current user.
	Any(abilities []string, arguments ...interface{}) bool

	// Authorize 确定是否应该为当前用户授予给定的能力
	// Determine if the given ability should be granted for the current user.
	Authorize(ability string, arguments ...interface{})

	// Inspect 给定能力的用户
	// the user for the given ability.
	Inspect(ability string, arguments ...interface{}) HttpResponse

	// ForUser 获取给定用户的警卫实例
	// Get a guard instance for the given user.
	ForUser(user Authorizable) Gate
}

type GateFactory interface {

	// Has 确定是否已定义给定的能力
	// determined if a given ability has been defined.
	Has(ability string) bool

	// Define 一种新的能力
	// a new ability.
	Define(ability string, callback GateChecker) GateFactory

	// Policy 为给定的类类型定义一个策略类
	// define a policy class for a given class type.
	Policy(class Class, policy Policy) GateFactory

	// Before 注册一个回调以在所有 Gate 检查之前运行
	// Register a callback to run before all Gate checks.
	Before(callable GateHook) GateFactory

	// After 注册回调以在所有 Gate 检查后运行
	// Register a callback to run after all Gate checks.
	After(callable GateHook) GateFactory

	// Check 确定是否应为当前用户授予所有给定的能力
	// Determine if all the given abilities should be granted for the current user.
	Check(user Authorizable, ability string, arguments ...interface{}) bool

	// Abilities 获得所有已定义的能力
	// Get all the defined abilities.
	Abilities() []string
}
