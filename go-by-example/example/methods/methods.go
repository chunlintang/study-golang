package main

import "fmt"

// 在结构体中定义方法
type rect struct {
	width, height int
}

// area方法接收器类型"rect""
func (r *rect) area() int {
	return r.width * r.height
}

// 值接收器
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// 使用指针调用方法，避免方法在调用时产生拷贝
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
