package _select

import (
	"fmt"
	"time"
)

// select 只能做用于管道读写
func SelectForChan(c chan string) {
	var recv string
	send := "Hello"

	select {
	case recv = <-c:
		fmt.Printf("received %s\n", recv)
	case c <- send:
		fmt.Printf("sent %s\n", send)
	}
}

// select 可以接受0~2个返回值
func SelectAssign(c chan string) {
	select {
	case <-c: // 0个变量
		fmt.Printf("0")
	case d := <-c: // 1个变量
		fmt.Printf("1: received %s\n", d)
	case d, ok := <-c: // 2个变量
		if !ok {
			fmt.Printf("no data found")
			break
		}
		fmt.Printf("2: received %s\n", d)
	}
}

func SelectDefault() {
	c := make(chan string)
	select {
	case <-c:
		fmt.Printf("received\n")
	default:
		fmt.Printf("no data found in default\n")
	}
}

func SelectChanStillBlock() {
	ch := make(chan int, 10)
	select {
	case _, ok := <-ch: // 此处仍然会阻塞
		fmt.Printf("never reached if channel is open. %v\n", ok)
	case <-time.After(10 * time.Millisecond):
		fmt.Printf("time out")
	}
}
