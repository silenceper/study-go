package main
import "fmt"
type Widuu int
func main() {
    var w Widuu
    w.G(100)
    fmt.Println(w)
}
func (a *Widuu) G(num int) {
    *a += Widuu(num)
}
