package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHttp")
	ctx := NewContext(request, writer)

	// 一个简单的路由选择器，这里直接写死为测试路由foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	router(ctx)

}

func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}


