### 九，Channel

```go
func channelDemo() {
	c := make(chan int) // 定义channel
	go func() {
		for {
			n := <-c // 接收数据
			fmt.Println(n)
		}
	}()
	c <- 1 // 发数据
	c <- 2
	time.Sleep(time.Minute)
}
```

- 通过通信来共享内存。