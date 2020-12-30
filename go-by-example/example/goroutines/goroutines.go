package main

import "fmt"

// 协程
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 同步
	f("direct")

	// 使用 go f(s) 在一个Go协程中调用这个函数
	// 这个新的Go协程会并发调用这个函数
	go f("goroutines")

	// 也可以使用匿名函数启动一个Go协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 上面启动了两个协程异步调用函数，程序直接运行到了这里，任意键结束程序
	fmt.Scanln()
	fmt.Println("done")
}
