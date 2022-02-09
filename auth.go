package contracts

type GuardDriver func(name string, config Fields, ctx Context, provider UserProvider) Guard
type UserProviderDriver func(config Fields) UserProvider

type Auth interface {
	ExtendUserProvider(name string, provider UserProviderDriver)
	ExtendGuard(name string, guard GuardDriver)

	Guard(name string, ctx Context) Guard
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


	Login(user Authenticatable) interface{}
}

type UserProvider interface {
	// RetrieveById 通过用户的唯一标识符检索用户
	// Retrieve a user by their unique identifier.
	RetrieveById(identifier string) Authenticatable
}

type Authorizable interface {
	Can(ability string, arguments ...interface{}) bool
}

type GateChecker func(user Authorizable, data ...interface{}) bool
type GateHook func(user Authorizable, ability string, data ...interface{}) bool

type Policy map[string]GateChecker

type Gate interface {

	// Allows determined if the given ability should be granted for the current user.
	Allows(ability string, arguments ...interface{}) bool

	// Denies Determine if the given ability should be denied for the current user.
	Denies(ability string, arguments ...interface{}) bool

	// Check Determine if all the given abilities should be granted for the current user.
	Check(abilities []string, arguments ...interface{}) bool

	// Any Determine if any one of the given abilities should be granted for the current user.
	Any(abilities []string, arguments ...interface{}) bool

	// Authorize Determine if the given ability should be granted for the current user.
	Authorize(ability string, arguments ...interface{})

	// Inspect the user for the given ability.
	Inspect(ability string, arguments ...interface{}) HttpResponse

	// ForUser Get a guard instance for the given user.
	ForUser(user Authorizable) Gate
}

type GateFactory interface {

	// Has determined if a given ability has been defined.
	Has(ability string) bool

	// Define a new ability.
	Define(ability string, callback GateChecker) GateFactory

	// Policy define a policy class for a given class type.
	Policy(class Class, policy Policy) GateFactory

	// Before Register a callback to run before all Gate checks.
	Before(callable GateHook) GateFactory

	// After Register a callback to run after all Gate checks.
	After(callable GateHook) GateFactory

	Check(user Authorizable, ability string, arguments ...interface{}) bool

	// Abilities Get all the defined abilities.
	Abilities() []string
}
