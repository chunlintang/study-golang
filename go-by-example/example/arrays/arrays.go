package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a) // [0 0 0 0 0]

	a[4] = 100
	fmt.Println(a) // [0 0 0 0 100]

	var b [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = i + j
		}
	}
	fmt.Println("array b:", b)
}
