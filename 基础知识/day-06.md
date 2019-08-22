### 四，面向接口

##### 一）定义

- 接口由使用者定义
- 接口的实现是隐式的

```go
// demo/main.go
package main

import (
	"fmt"
	"test/learn/day-07/retriever/mock"
	"time"
)

// 使用者定义接口
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
  // 接口变量自带指针
	r = &mock.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Second,
	}
	fmt.Println(download(r))
}

```

```go
// demo/mock/retriever.go
package mock

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}

// 接口实现
func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	ret, err := httputil.DumpResponse(resp, true)
	_ = resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(ret)
}
```

- 接口变量自带指针
- 接口变量同样采用值传递，几乎不需要使用接口指针
- 指针接收者实现只能以指针方式使用；值接收者都可以

##### 二）查看接口类型

- Type assert
- Type Switch