package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Join:", strings.Join([]string{"a", "b", "c", "d"}, "-")) // a-b-c-d
	fmt.Println("Replace:", strings.Replace("foo", "o", "0", -1))         // f00
	fmt.Println("Replace:", strings.Replace("foo", "o", "0", 1))          // f0o
	fmt.Println("Split:", strings.Split("a-b-c-d", "-"))                  // [a b c d]
}
