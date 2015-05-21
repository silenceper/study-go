package main

import "fmt"

type Rect struct{
    x,y float64
    width,height float64
}

func (r *Rect) Area() float64{
    r.width=400
    return r.width*r.height 
}

func main(){
    /* 初始化的方法
    rect1 := new(Rect)
    rect2 := &Rect{}
    rect3 := &Rect{0, 0, 100, 200}
    rect4 := &Rect{width: 100, height: 200}
    */
    rect := Rect{0,0,100,200}
    fmt.Println(rect.Area())
    fmt.Println(rect.width)
}
