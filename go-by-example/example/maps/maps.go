package main

import "fmt"

func main() {
	// 要创建一个空 map，需要使用内建函数 make：make(map[key-type]val-type)。
	m := make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2

	fmt.Println("map:", m)

	value1 := m["key1"]
	fmt.Println("value1:", value1)

	fmt.Println("m length:", len(m))

	// 删除
	delete(m, "key2")
	fmt.Println("delete after m:", m)

	_, prs := m["key2"]
	fmt.Println("prs:", prs) // false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n) // map[bar:2 foo:1]
}
