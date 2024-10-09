package contracts

type Views interface {
	Render(name string, data ...any) HttpResponse
	Register(name, template string)
}
