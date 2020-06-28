package _range

import (
	"fmt"
	"sync"
	"time"
)

// 题目一：性能考察
func FindMonkey(s []string) bool {
	for _, v := range s {
		if v == "monkey" {
			return true
		}
	}

	return false
}

func FindMonkeyImproved(s []string) bool {
	for i := range s {
		if s[i] == "monkey" {
			return true
		}
	}

	return false
}

// 题目二： 引用循环变量下标
func PrintSlice() {
	s := []int{1, 2, 3}
	var wg sync.WaitGroup

	wg.Add(len(s))
	for _, v := range s {
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}

// 题目三：range nil channel 永久阻塞
func RangeNilChannel() {
	var c chan string

	for v := range c {
		fmt.Println(v)
	}
}

/*
- A： 函数没有打印，正常退出
- B： 遍历nil切片，函数会panic
- C： 函数打印nil，然后退出
- D： 函数没有打印，永远阻塞
*/

// 题目四： range 会一直阻塞到channel关闭
func RangeTimer() {
	t := time.NewTimer(time.Second)

	for _ = range t.C {
		fmt.Println("hi")
	}
}

/*
- A： 函数没有打印，直接退出
- B： 函数打印"hi"后退出
- C： 函数打印"hi"后阻塞
- D： 函数panic
*/

// 题目五：动态遍历
func RangeDemo() {
	s := []int{1, 2, 3}
	for i := range s {
		s = append(s, i)
	}
}
