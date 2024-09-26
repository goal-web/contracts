package contracts

import (
	"net/url"
)

type RouteParams map[string]string

type Route interface {
	// Middlewares 获取附加到路由的中间件
	// Get the middlewares attached to the route.
	Middlewares() []MagicalFunc

	// Method 获取附加到路由的请求方法
	// Get the request method attached to the route.
	Method() []string

	// GetPath 获取附加到路由的请求路径
	// Get the request path attached to the route.
	GetPath() string
	GetHost() string
	Host(host string) Route
	GetName() string

	// Handler 获取附加到路由的路由处理程序
	// Get the route handler attached to the route.
	Handler() MagicalFunc

	Name(name string) Route
}

type RouteGroup interface {
	// Get 向路由组注册一个新的 GET 路由。
	// Register a new GET route with the route group.
	Get(path string, handler any, middlewares ...any) RouteGroup
	GET(path string, handler any, middlewares ...any) RouteGroup

	// Post 向路由组注册一个新的 POST 路由。
	// Register a new POST route with the route group.
	Post(path string, handler any, middlewares ...any) RouteGroup
	POST(path string, handler any, middlewares ...any) RouteGroup

	// Delete 向路由组注册一个新的 DELETE 路由。
	// Register a new DELETE route with the routing group.
	Delete(path string, handler any, middlewares ...any) RouteGroup
	DELETE(path string, handler any, middlewares ...any) RouteGroup

	// Put 向路由组注册一个新的 PUT 路由。
	// Register a new PUT route with the routing group.
	Put(path string, handler any, middlewares ...any) RouteGroup
	PUT(path string, handler any, middlewares ...any) RouteGroup

	// Patch 向路由组注册一个新的 PATCH 路由。
	// Register a new PATCH route with the routing group.
	Patch(path string, handler any, middlewares ...any) RouteGroup
	PATCH(path string, handler any, middlewares ...any) RouteGroup

	// Options 向路由组注册一个新的 OPTIONS 路由。
	// Register a new OPTIONS route with the routing group.
	Options(path string, handler any, middlewares ...any) RouteGroup
	OPTIONS(path string, handler any, middlewares ...any) RouteGroup

	// Trace 向路由组注册一个新的 TRACE 路由。
	// Register a new TRACE route with the routing group.
	Trace(path string, handler any, middlewares ...any) RouteGroup
	TRACE(path string, handler any, middlewares ...any) RouteGroup

	// Group 创建具有共享属性的路由组。
	// Create a route group with shared attributes.
	Group(prefix string, middlewares ...any) RouteGroup

	GetHost() string
	Host(host string) RouteGroup

	// Routes 获取路由
	// get route.
	Routes() []Route
}

type HttpRouter interface {
	// Get 向路由器注册一个新的 GET 路由。
	// Register a new GET route with the router.
	Get(path string, handler any, middlewares ...any) Route
	GET(path string, handler any, middlewares ...any) Route

	// Post 向路由器注册一个新的 POST 路由。
	// Register a new POST route with the router.
	Post(path string, handler any, middlewares ...any) Route
	POST(path string, handler any, middlewares ...any) Route

	// Delete 向路由器注册一个新的 DELETE 路由。
	// Register a new DELETE route with the router.
	Delete(path string, handler any, middlewares ...any) Route
	DELETE(path string, handler any, middlewares ...any) Route

	// Put 向路由器注册一个新的 PUT 路由。
	// Register a new PUT route with the router.
	Put(path string, handler any, middlewares ...any) Route
	PUT(path string, handler any, middlewares ...any) Route

	// Patch 向路由器注册一个新的 PATCH 路由。
	// Register a new PATCH route with the router.
	Patch(path string, handler any, middlewares ...any) Route
	PATCH(path string, handler any, middlewares ...any) Route

	// Options 向路由器注册一个新的 OPTIONS 路由。
	// Register a new OPTIONS route with the router.
	Options(path string, handler any, middlewares ...any) Route
	OPTIONS(path string, handler any, middlewares ...any) Route

	// Trace 向路由器注册一个新的 TRACE 路由
	// Register a new TRACE route with the router.
	Trace(path string, handler any, middlewares ...any) Route
	TRACE(path string, handler any, middlewares ...any) Route

	// Use  使用中间件
	// use middleware.
	Use(middlewares ...any)

	// Middlewares 返回全局中间件
	Middlewares() []MagicalFunc

	// Group 创建具有共享属性的路由组。
	// Create a route group with shared attributes.
	Group(prefix string, middlewares ...any) RouteGroup

	// Mount 装配路由
	Mount() error

	// Route 通过 url 找到合适的路由
	Route(method string, url *url.URL) (Route, RouteParams, error)

	// Print 打印所有路由
	Print()
}

type Router[T any] interface {
	Find(route string) (T, RouteParams, error)
	Add(route string, data T) (string, error)
	IsEmpty() bool
	All() []T
}
