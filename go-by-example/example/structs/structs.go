package main

import "fmt"

// 结构体是带类型的字段集合，组织数据非常有用

type person struct {
	name string
	age  int
}

func main() {
	// 创建新的结构体元素
	fmt.Println(person{"Jack", 20})

	// 初始化元素指定字段
	fmt.Println(person{name: "Tom", age: 22})

	// 省略的字段将被初始化为零值
	fmt.Println(person{name: "Bob"})

	// 结构体指针
	fmt.Println(&person{name: "Jery", age: 18})

	// 访问结构体字段
	s := person{name: "Mantis", age: 22}
	fmt.Println(s.name)

	// 也可对结构体指针使用"." 指针会自动去解引用
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
