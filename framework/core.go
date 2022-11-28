package framework

import (
	"net/http"
)

type Core struct {

}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// TODO
}

func NewCore() *Core {
	return &Core{}
}


