### 八，Goroutine

```go
func main() {
	for i := 0; i < 100000; i++ {
		go func(j int) {
			// 并发执行
			for {
				fmt.Printf("Hello from "+"goroutine %d\n", j)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
```

##### 一）协程Cotoutine

- 轻量级“线程”
- 非抢占式多任务处理，由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或者多个线程上运行
- 子程序是协程的一个特列

```go
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(j int) {
			// 并发执行
			for {
				a[j]++ // 无法主动交出控制权
				runtime.Gosched() // 手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
```

##### 二）调度器

- 任何函数只需要加上go就能送给调度器运行
- 不需要在定义时区分是否是异步调用
- 调度器在合适的点进行切换
- 使用-race来检测数据访问冲突

##### 三）goroutine可能的切换点

- I/O、select
- channel
- 等待锁
- 只是参考，不能保证切换，不能保证其他地方不切换