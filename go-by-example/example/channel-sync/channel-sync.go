package main

import (
	"fmt"
	"time"
)

// 通道同步
// 使用通道同步Go协程间的执行状态，用一个阻塞的接收方式来等待一个Go协程的结束

// 在协程中运行函数,done
// 通知其他协程已经完成
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值通知已经完成
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
