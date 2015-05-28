package main

import (
    _ "./providers/memory/"
    //"github.com/astaxie/session"
    "./session/"
    "fmt"
)

var globalSessions *session.Manager
//然后在init函数中初始化
func init() {
    globalSessions, _ = session.NewManager("memory","gosessionid",3600)
    
    go globalSessions.GC()
}

func main(){
    fmt.Println(globalSessions)
}
