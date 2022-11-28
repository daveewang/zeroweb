package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"zeroweb/framework"
)

func FooControllerHandler(ctx *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	// 浏览器的问题，用浏览器请求，会先请求/foo，再请求/favicon.ico。
	//检测方案： 可以在/controller中打印出url检查。解决方案：使用postman或者别的客户端发送请求
	fmt.Println(ctx.GetRequest().URL)

	go func() {
		defer func() {
			if p:=recover();p!=nil{
				panicChan <- p
			}
		}()
		time.Sleep(10*time.Second)
		ctx.Json(200, "OK")
		finish <- struct{}{}
	}()

	select {
	case p:=<-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Println(p)
		ctx.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(500, "time out")
		ctx.SetHasTimeout()
	}

	return nil
}
