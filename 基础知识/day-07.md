### 五，函数式编程

- 函数是一等公民：参数、变量、返回值都可以是函数
- 高阶函数
- 函数->闭包

```go
package main

import "fmt"

// 函数式编程、闭包
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// 正统函数式编程
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	b := adder()
	for j := 0; j <= 10; j++ {
		fmt.Printf("0+1+2+....+%d=%d\n", j, b(j))
	}

	a := adder2(0)
	for i := 0; i <= 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+1+2+....+%d=%d\n", i, s)
	}
}
```

