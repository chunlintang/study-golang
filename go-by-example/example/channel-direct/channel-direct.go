package main

import "fmt"

// 当使用通道时，可以指定这个通道是不是只用来发送或者接收值

// ping只用来发送值
func ping(pings chan<- string, msg string) {

	pings <- msg
}

// pong只用来接收值
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
