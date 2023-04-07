package contracts

type Route interface {
	// Middlewares 获取附加到路由的中间件
	// Get the middlewares attached to the route.
	Middlewares() []MagicalFunc

	// Method 获取附加到路由的请求方法
	// Get the request method attached to the route.
	Method() []string

	// Path 获取附加到路由的请求路径
	// Get the request path attached to the route.
	Path() string

	// Handler 获取附加到路由的路由处理程序
	// Get the route handler attached to the route.
	Handler() MagicalFunc
}

type RouteGroup interface {
	// Get 向路由组注册一个新的 GET 路由。
	// Register a new GET route with the route group.
	Get(path string, handler any, middlewares ...any) RouteGroup

	// Post 向路由组注册一个新的 POST 路由。
	// Register a new POST route with the route group.
	Post(path string, handler any, middlewares ...any) RouteGroup

	// Delete 向路由组注册一个新的 DELETE 路由。
	// Register a new DELETE route with the routing group.
	Delete(path string, handler any, middlewares ...any) RouteGroup

	// Put 向路由组注册一个新的 PUT 路由。
	// Register a new PUT route with the routing group.
	Put(path string, handler any, middlewares ...any) RouteGroup

	// Patch 向路由组注册一个新的 PATCH 路由。
	// Register a new PATCH route with the routing group.
	Patch(path string, handler any, middlewares ...any) RouteGroup

	// Options 向路由组注册一个新的 OPTIONS 路由。
	// Register a new OPTIONS route with the routing group.
	Options(path string, handler any, middlewares ...any) RouteGroup

	// Trace 向路由组注册一个新的 TRACE 路由。
	// Register a new TRACE route with the routing group.
	Trace(path string, handler any, middlewares ...any) RouteGroup

	// Middlewares 获取附加到路由的中间件
	// Get the middlewares attached to the route.
	Middlewares() []MagicalFunc

	// Group 创建具有共享属性的路由组。
	// Create a route group with shared attributes.
	Group(prefix string, middlewares ...any) RouteGroup

	// Routes 获取路由
	// get route.
	Routes() []Route

	// Groups 获取路由组
	// get routing group.
	Groups() []RouteGroup
}

type Router interface {
	Static(path string, directory string)
	// Get 向路由器注册一个新的 GET 路由。
	// Register a new GET route with the router.
	Get(path string, handler any, middlewares ...any)

	// Post 向路由器注册一个新的 POST 路由。
	// Register a new POST route with the router.
	Post(path string, handler any, middlewares ...any)

	// Delete 向路由器注册一个新的 DELETE 路由。
	// Register a new DELETE route with the router.
	Delete(path string, handler any, middlewares ...any)

	// Put 向路由器注册一个新的 PUT 路由。
	// Register a new PUT route with the router.
	Put(path string, handler any, middlewares ...any)

	// Patch 向路由器注册一个新的 PATCH 路由。
	// Register a new PATCH route with the router.
	Patch(path string, handler any, middlewares ...any)

	// Options 向路由器注册一个新的 OPTIONS 路由。
	// Register a new OPTIONS route with the router.
	Options(path string, handler any, middlewares ...any)

	// Trace 向路由器注册一个新的 TRACE 路由
	// Register a new TRACE route with the router.
	Trace(path string, handler any, middlewares ...any)

	// Use  使用中间件
	// use middleware.
	Use(middlewares ...any)

	// Group 创建具有共享属性的路由组。
	// Create a route group with shared attributes.
	Group(prefix string, middlewares ...any) RouteGroup

	// Start 启动 httpserver
	// start httpserver.
	Start(address string) error

	// Close 关闭 httpserver
	// close httpserver.
	Close() error
}
