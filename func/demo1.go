package main
import "fmt"
var (
    i int = 10
)
var j int = 20

func add(a int,b int)int{
    return a+b
}

func main(){
    sum :=add(i,j)
    fmt.Printf("%d+%d=%d\n",i,j,sum)
}

