package main

import "fmt"

// 通道缓冲
// 默认通道是没有缓冲的，只有对应的接收通道准备好接收时，才允许进行发送
// 缓存通道允许在没有接收方的情况下，缓存限定数量的值
func main() {
	// 创建一个通道，最多缓存2个值
	messages := make(chan string, 2)

	// 缓冲通道可以直接将值发送到通道中
	messages <- "buffered"
	messages <- "channel"
	//messages <- "channel2" // 报错，最多缓存2个

	// 接收值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
