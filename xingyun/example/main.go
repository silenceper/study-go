package main

import (
	"net/http"
    "./test"
	"github.com/silenceper/xingyun"
    "fmt"
)

var logger xingyun.Logger

func main() {
	cfg := &xingyun.Config{}
	server := xingyun.NewServer(cfg)
	logger = server.Logger()
    test.Name="silenceper"
    
	pipe := server.NewPipe("test", xingyun.PipeHandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		next(w, r)
	}))

	server.Get("/hello", pipe.Wrap(func(ctx *xingyun.Context) {
		ctx.SetSession("test",[]byte("yes"))
        ctx.WriteString("hello world")
    }))

	server.Get("/halt", pipe.Wrap(func(ctx *xingyun.Context) {
		fmt.Println(ctx.GetSession("test"))
        ctx.WriteString("halt")
	}))

	err := server.ListenAndServe("0.0.0.0:9000")
	logger.Errorf("%s", err)
}
