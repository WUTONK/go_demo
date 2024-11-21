package main

import "fmt"

// 定义一个接口
type Shape interface {
	Area() float64
}

// 实现接口的结构体
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 使用接口进行多态
func PrintArea(s Shape) {
	fmt.Printf("面积: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	PrintArea(c) // 输出: 面积: 78.50
	PrintArea(r) // 输出: 面积: 12.00
}
