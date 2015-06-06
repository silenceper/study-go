package main 
import(
	"math"
	"fmt"
)
func main() {
	var (
		count =100.0
		limit =13.0
	)
	size:=math.Ceil(count/limit)

	fmt.Println(size,int64(size))

}