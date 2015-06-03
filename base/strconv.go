package main
import(
    "strconv"
    "fmt"
)

func main(){
    var a string="0"
    var b string="1aa"
    str1,err:=strconv.Atoi(a)
    str2,err:=strconv.Atoi(b)
    if err!=nil{
        panic(err)
    }
    fmt.Println(str1,str2)
}
