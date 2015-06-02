package main

import(
    "time"
    "fmt"
)

func main(){
    var t int64 = time.Now().Unix()
    var s string = time.Unix(t, 0).Format("2006-01-02 15:04:05")
    fmt.Println(s)
}
