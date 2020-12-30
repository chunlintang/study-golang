package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for {
		fmt.Println("break")
		break
	}
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			fmt.Println("continue")
			continue
		}
		fmt.Println(n)
	}
}
