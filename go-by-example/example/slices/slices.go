package main

import "fmt"

func main() {
	// 要创建一个长度不为 0 的空 slice，需要使用内建函数 make
	s := make([]string, 3)
	fmt.Println("s:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("s:", s)

	fmt.Println("s len:", len(s))

	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f")

	c := make([]string, len(s))
	copy(c, s) // 拷贝
	fmt.Println("c copy from s:", c)

	l := s[2:5]
	fmt.Println("s[2:5]:", l)
	fmt.Println("s[:5]:", s[:5])
	fmt.Println("s[2:]:", s[2:])
}
