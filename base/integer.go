package main
import (
    "fmt"
    "strconv"
)

func main(){
    var a int64=123123123
    str:=strconv.FormatInt(a,10)
    fmt.Println(a,str)
}
