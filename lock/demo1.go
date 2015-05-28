package main
 
import (
    "fmt"
    "sync"
    "time"
)
 
var m *sync.Mutex
 
func main() {
    m = new(sync.Mutex)
 
    go lock(1)
    time.Sleep(time.Second)
     
    lock(2)
 
    fmt.Printf("%s\n", "exit!")
}
 
func lock(i int){
    println(i, "lock start")
 
    m.Lock()
    println(i, "lock")
 
    time.Sleep(10 * time.Second)
 
    m.Unlock()
    println(i, "unlock")
}
