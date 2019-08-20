### 二，内建容器

##### 一）数组

- 定义数组

```go
var arr1 [5]int
arr2 := [3]int{1, 2, 4}
arr3 := [...]int{2, 4, 6, 7, 8, 9}
var grid [4][5]int // 4行5列，默认数组元素都是0,[[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
```

- 遍历数组

```go
// 第一种
for i, v := range arr3 {
  fmt.Println(i, v)
}
// 第二种
for i := 0;i < len(grid);i++ {
  fmt.Println(grid[i])
}
```

- 数组是值类型
  - [10]int和[20]int是不同的类型
  - func foo(arr [10]int)会拷贝数组
  - 在Go中我们一般不使用数组，而会使用切片

```go
func printArr(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	arr2 := [...]int{3, 4, 5, 6, 7}
	printArr(arr2)
  //0 100
	//1 4
	//2 5
	//3 6
	//4 7
	fmt.Println(arr2)
  //[3 4 5 6 7] 
}
```

使用指针

```go
func printArr(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func main() {
	arr2 := [...]int{3, 4, 5, 6, 7}
	printArr(&arr2)
  //0 100
	//1 4
	//2 5
	//3 6
	//4 7
	fmt.Println(arr2)
  //[100 4 5 6 7] 
}
```

##### 二）切片（Slice）

- 定义

数组：[5]int

切片：[]int

```go
func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
}
//arr[2:6] =  [2 3 4 5]
//arr[:6] =  [0 1 2 3 4 5]
//arr[2:] =  [2 3 4 5 6 7]
//arr[:] =  [0 1 2 3 4 5 6 7]
```

- 切片是按引用传递

```go
func updateSlice(s []int)  {
	s[0] = 100
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}

	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)
	fmt.Println("After update slice")
	updateSlice(s1)
	fmt.Println(s1)
  fmt.Println(s2)
	fmt.Println(arr)
}
//s1 =  [2 3 4 5 6 7]
//s2 =  [0 1 2 3 4 5 6 7]
//After update slice
//[100 3 4 5 6 7]
//[0 1 100 3 4 5 6 7]
//[0 1 100 3 4 5 6 7]
```

- 切片分割后，它仍然有原始切片的隐藏下标

```go
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)
}
//[2 3 4 5]
//[5 6]
```

![](https://i.loli.net/2019/08/11/KMsy3jvRbYFlVrD.png)

切片的底层数据结构

![](https://i.loli.net/2019/08/11/aTkXVcsqRuwS4Of.jpg)

- slice的扩展
  - slice可以向后扩展，不可以向前扩展
  - s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)

```go
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)
  
  fmt.Println("============================")
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))
  
  fmt.Println(s1[3:7]) // 越界
}
//[2 3 4 5]
//[5 6]
============================
//s1=[2 3 4 5], len(s1)=4, cap(s1)=6
//s2=[5 6], len(s2)=2, cap(s2)=3

// 越界
//panic: runtime error: slice bounds out of range
```

