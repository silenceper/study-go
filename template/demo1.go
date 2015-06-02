package main

import (
    "net/http"
    "path"
    "fmt"
    "html/template"
)
type Person struct{
    UserName string
    Age int32
}

func testFunc(arg string)string{
    return "test"
}

func handleMain(w http.ResponseWriter,r *http.Request){
    funcMap:=template.FuncMap{"testFunc":testFunc} 
    var tpl="./tmpl/first.html"
    //使用path.Base 获取文件名 必须这么使用
    t, err := template.New(path.Base(tpl)).Funcs(funcMap).ParseFiles(tpl)
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
