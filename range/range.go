package _range

import (
	"fmt"
	"time"
)

func ForExpression() {
	s := []int{1, 2, 3}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func ForRangeExpression() {
	s := []int{1, 2, 3}

	for i := range s {
		fmt.Println(s[i])
	}
}

// RangeArray 演示Range遍历数组
func RangeArray() {
	a := [3]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("index: %d value:%d\n", i, v)
	}
}

// RangeArrayPointer 演示range遍历数组的指针
// 数组的指针与数组一致
func RangeArrayPointer() {
	a := [3]int{1, 2, 3}
	for i, v := range &a {
		fmt.Printf("index: %d value:%d\n", i, v)
	}
}

// Range遍历切片与遍历数组操作一致。
// 但不能作用于切片的地址，否则会有编译错误：
// cannot range over &s (type *[]int)
func RangeSlice() {
	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Printf("index: %d value:%d\n", i, v)
	}
}

func RangeString() {
	s := "Hello"
	for i, v := range s {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}
}

// len(“中国”) == 6，长度代表底层存储
func RangeStringUniCode() {
	s := "中国"
	for i, v := range s {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}
}

func RangeMap() {
	m := map[string]string{"animal": "monkey", "fruit": "apple"}
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func RangeChannel() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	time.AfterFunc(time.Microsecond, func() {
		close(c)
	})

	for e := range c {
		fmt.Printf("element: %s\n", e)
	}
}
