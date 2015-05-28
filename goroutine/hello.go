package main
import (
    "fmt"
    //"runtime"
    "time"
)

func say(s string){
    for i:=0;i<5;i++{
        //runtime.Gosched()
        fmt.Println(s)
    }
}

func main(){
    go say("World")
    //say("Hello")
    time.Sleep(time.Second*1)
}
