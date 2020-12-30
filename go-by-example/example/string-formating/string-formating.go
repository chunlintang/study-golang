package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)    // {1 2}
	fmt.Printf("%+v\n", p)   // {x:1 y:2}
	fmt.Printf("%#v\n", p)   // main.point{x:1, y:2}
	fmt.Printf("%t\n", true) // true
	fmt.Printf("%d\n", 123)  // 123 十进制格式化
	fmt.Printf("%b\n", 14)   // 1110 二进制表示
}
