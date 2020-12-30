package main

import (
	"errors"
	"fmt"
)

// 错误处理
// 内建接口error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("error args 42")
	}

	return arg + 3, nil
}

// 自定义error类型
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "error args 42"}
	}

	return arg + 3, nil
}
