package main

import (
    "net/http"
    //"io"
    "fmt"
    "html/template"
)
type Person struct{
    UserName string
    Age int32
}

func handleMain(w http.ResponseWriter,r *http.Request){
    //指定一个模板  不需要New
    t, err := template.ParseFiles("./tmpl/first.html")
    if err!=nil {
        fmt.Println(err)
        return
    }
    p:=Person{"wenzhenlin",23}
    t.Execute(w, p)
}

func main(){
    http.HandleFunc("/",handleMain)
    err := http.ListenAndServe(":9000",nil)
    if err!=nil{
        fmt.Println(err)
    }
}
