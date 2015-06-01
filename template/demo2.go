package main
import (
    "net/http"
    "html/template"
    //"fmt"
)

func handleMain(w http.ResponseWriter,r *http.Request){
    t:=template.New("test_template")
    t,_=t.Parse("hello world")
    t.Execute(w,nil)
}

func main(){
    http.HandleFunc("/",handleMain)
    http.ListenAndServe(":9000",nil)
}
