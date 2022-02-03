package contracts

type GuardDriver func(config Fields, ctx Context, provider UserProvider) Guard
type UserProviderDriver func(config Fields) UserProvider

type Auth interface {
	ExtendUserProvider(name string, provider UserProviderDriver)
	ExtendGuard(name string, guard GuardDriver)

	Guard(name string, ctx Context) Guard
	UserProvider(name string) UserProvider
}

type Authenticatable interface {
	GetId() string
}

type Guard interface {
	Once(user Authenticatable)
	User() Authenticatable
	GetId() string
	Check() bool
	Guest() bool
	Validate(credentials Fields) bool
	Login(user Authenticatable) interface{}
}

type UserProvider interface {
	RetrieveById(identifier string) Authenticatable
}
