package main

import "fmt"

// 通道，它是连接多个协程的管道，可以从一个Go协程将值发送到通道，然后在别的Go协程中接收
func main() {

	// 使用make(chan value-type)创建一个通道
	messages := make(chan string)

	// 使用 channel <- 发送一个值到通道中
	// 发送"ping"到messages通道中
	go func() {
		messages <- "ping"
	}()

	// 使用 <-channel 从通道中接收一个值
	msg := <-messages
	fmt.Println(msg)

	// 默认发送和接收操作都是阻塞的，知道发送方和接收方都准备完毕，这个新允许我们不使用其他的同步操作，来在程序结尾等待消息
}
