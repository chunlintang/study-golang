package main

import (
	"fmt"
	"sort"
)

func main() {
	// 字符
	str := []string{"c", "f", "a", "d", "g", "b"}
	strSorted := sort.StringsAreSorted(str)
	fmt.Println("sorted str:", strSorted)
	if strSorted == false {
		sort.Strings(str)
		fmt.Println("Strings stort:", str)
	}
	// 整数
	ints := []int{3, 2, 10, 5, 8, 1}
	intSorted := sort.IntsAreSorted(ints)
	fmt.Println("soted int:", intSorted)
	if intSorted == false {
		sort.Ints(ints)
		fmt.Println("int sort:", ints)
	}
}
