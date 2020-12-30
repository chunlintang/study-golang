package main

import "fmt"

// 使用一个带default子句的select来实现非阻塞发送、接收

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 接收
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message default")
	}

	// 发送
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message send")
	}

	// 接收多个
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
