package main

import (
	"fmt"
	"time"
)

// 线程池
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " processing job")
		time.Sleep(time.Second * 1)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 分成三个worker，初始化是阻塞的，未传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 15个任务,每个worker执行5个job
	for j := 1; j <= 15; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 15; a++ {
		<-results
	}
}
