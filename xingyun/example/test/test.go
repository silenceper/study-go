package test

import "fmt"

var (
    Name string = ""
)

func init(){
    fmt.Println("init test")
}

func Testlog(){
    fmt.Printf(Name)
    Name="wenzhenlin"
}
