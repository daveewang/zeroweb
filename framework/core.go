package framework

import (
	"log"
	"net/http"
	"strings"
)

// 框架核心架构
type Core struct {
	router map[string]*Tree
}


// 路由注册  Get、Post、Put、Delete
func (c *Core) Get(url string, handler ControllerHandler) {
	err := c.router["GET"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router err: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	err := c.router["POST"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router err: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	err := c.router["PUT"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router err: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	err := c.router["DELETE"].AddRouter(url, handler)
	if err != nil {
		log.Fatal("add router err: ", err)
	}
}

// 匹配路由，如果没有匹配到，则返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层
	if methodHandlers, ok:=c.router[upperMethod];ok{
		return methodHandlers.FindHandler(uri)
	}
	return nil

}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 所有请求都会在这处理
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHttp")
	// 封装自定义context
	ctx := NewContext(request, response)

	// 寻找路由
	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}

	// 调用路由，如果返回err代表内部存在错误，返回500
	err := router(ctx)
	if err != nil {
		ctx.Json(500, "inner error")
		return
	}

}

// 初始化核心架构
func NewCore() *Core {
	// 将二级map写入一级map中
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}
}


