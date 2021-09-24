package contracts

import "github.com/labstack/echo/v4"

// HttpResponse http 响应
type HttpResponse interface {
	Status() int
	Response(ctx echo.Context) error
}

// HttpRequest http 请求
type HttpRequest interface {
	echo.Context
	All() Fields
}
