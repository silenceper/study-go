package main
import(
    "path"
    "fmt"
)

func main(){
    dir:="/root/workspaces/"
    file:="text.txt"
    fmt.Println(path.Join(dir,file))

}
