package main

import (
	"net/http"
    "./test"
    "./test2"
	"github.com/xiaoenai/xingyun"
)

var logger xingyun.Logger

func main() {
	cfg := &xingyun.Config{}
	server := xingyun.NewServer(cfg)
	logger = server.Logger()
    test.Name="silenceper"
    
	pipe := server.NewPipe("test", xingyun.PipeHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		logger.Tracef("before")
		next(w, r)
		logger.Tracef("after")
	}))

	server.Get("/hello", pipe.Wrap(func(ctx *xingyun.Context) {
		//ctx.Halt()
		logger.Tracef("hello world")
	    test.Testlog()
        test2.Test2()
		ctx.WriteString("hello world")
    }))

	server.Get("/halt", pipe.Wrap(func(ctx *xingyun.Context) {
		//ctx.Halt()
		logger.Tracef("halt")
		ctx.WriteString("halt")
	}))

	err := server.ListenAndServe("0.0.0.0:9000")
	logger.Errorf("%s", err)
}
