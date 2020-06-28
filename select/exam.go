package _select

import (
	"fmt"
)

/*
下面函数输出什么？
单选：
- A： 函数输出“c1”
- B： 函数输出“c2”
- C： 函数输出“c1”、“c2”
- D： 函数可能输出“c1”，也可能输出"c2"
*/
func SelectExam1() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	c1 <- 1
	c2 <- 2

	select {
	case <-c1:
		fmt.Println("c1")
	case <-c2:
		fmt.Println("c2")
	}
}

/*
下面函数输出什么？
单选：
- A： 函数输出“readable”
- B： 函数输出"writable"
- C： 函数什么也不输出，正常返回
- D： 函数什么也不输出，陷入阻塞
*/
func SelectExam2() {
	c := make(chan int)

	select {
	case <-c:
		fmt.Printf("readable")
	case c <- 1:
		fmt.Println("writable")
	}
}

/*
下面函数输出什么？
单选：
- A： 函数输出"1"
- B： 函数输出空字符
- C： 函数什么也不输出，陷入阻塞
- D： 函数什么也不输出，发生panic
*/
func SelectExam3() {
	c := make(chan int, 10)
	c <- 1
	close(c)

	select {
	case d := <-c:
		fmt.Println(d)
	}
}

/*
下面函数输出什么？
单选：
- A： 函数输出“1”
- B： 函数输出“no data received”
- C： 函数什么也不输出，陷入阻塞
- D： 函数什么也不输出，发生panic
*/
func SelectExam4() {
	c := make(chan int, 10)
	c <- 1
	close(c)

	select {
	case d, ok := <-c:
		if !ok {
			fmt.Println("no data received")
			break
		}
		fmt.Println(d)
	}
}

/*
关于下面函数的描述，正确的是？
单选：
- A： 编译错误，select语句非法
- B： 运行时错误，panic
- C： 函数陷入阻塞
- D： 函数什么也不做，直接返回
*/
func SelectExam5() {
	select {}
}

/*
关于下面函数的描述，正确的是？
单选：
- A： 函数会因为写值为nil的管道而被阻塞
- B： 函数会因为写值为nil的管道而panic
- C： 函数会从default出口返回
- D： 编译错误，值为nil的管道不可以写
*/
func SelectExam6() {
	var c chan string
	select {
	case c <- "Hello":
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}
