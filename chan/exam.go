package _chan

import (
	"fmt"
	"sync"
)

/*
以下可以实现互斥锁的是？
- A：
```
var counter int = 0
var ch = make(chan int, 1)

func Worker() {
	ch <- 1
	counter++
	<-ch
}
```

- B：
```go
var counter int = 0
var ch = make(chan int)

func Worker() {
	<-ch
	counter++
	ch <- 1
}
```

- C：
```go
var counter int = 0
var ch = make(chan int, 1)

func Worker() {
	<-ch
	counter++
	ch <- 1
}
```

- D：
```go
var counter int = 0
var ch = make(chan int)

func Worker() {
	ch <- 1
	counter++
	<-ch
}
```

*/
var counter int = 0
var ch = make(chan int, 1)
var wg sync.WaitGroup

func Worker() {
	ch <- 1
	counter++
	wg.Done()
	fmt.Println(counter)
	<-ch
}

/*
下面关于管道的描述正确的是？
单选：
- A： 读nil管道会panic
- B： 写nil管道会panic
- C： 读关闭的管道会panic
- D： 写关闭的管道会panic

*/

// 读nil管道会阻塞
func ChanReadNil() {
	var ch chan int
	d := <-ch
	fmt.Println(d)
}

// 写nil管道不会阻塞
func ChanWriteNil() {
	var ch chan int
	ch <- 1
}

// 关闭的管道可以读
func ChanReadClosed() {
	ch := make(chan int, 10)
	close(ch)
	d := <-ch
	fmt.Println(d)
}

// 关闭的管道不可写，否则panic
// send on closed channel
func ChanWriteClosed() {
	ch := make(chan int, 10)
	close(ch)
	ch <- 1
}

/*
下面函数输出什么？

*/

func ChanCap() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
