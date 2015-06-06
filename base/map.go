package main

import "fmt"

func main() {
	var data=make(map[string]interface{})
	var arr []string
	//var str string
	//data=nil

	fmt.Println(data==nil)
	fmt.Println(arr==nil)
	//fmt.Println(str==nil)  //mis match error
	data["dd"]="dd"

	fmt.Println("Hello, 世界")
}
