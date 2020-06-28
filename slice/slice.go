package _slice

import "fmt"

func SliceDeclare() {
	var s []int
	fmt.Println(s == nil) // true
}

func SliceLiteral() {
	s1 := []int{}          // 空切片
	s2 := []int{1, 2, 3}   // 长度为3的切片
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false
}

// make不指定空间时，空间默认等于长度
func SliceInitMake() {
	s1 := make([]int, 12)      // 指定长度
	s2 := make([]int, 10, 100) // 指定长度和空间

	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false
}

func SliceCut() {
	array := [5]int{1, 2, 3, 4, 5}
	s1 := array[0:2] // 从数组切取
	s2 := s1[0:1]    // 从切片切取
	fmt.Println(s1)  // [1 2]
	fmt.Println(s2)  // [1]
}

func SliceInitNew() {
	s := *new([]int)
	fmt.Println(s == nil) // true
}

func SliceAppend() {
	s := make([]int, 0)
	s = append(s, 1)              // 添加1个元素
	s = append(s, 2, 3, 4)        // 添加多个元素
	s = append(s, []int{5, 6}...) // 添加一个切片
	fmt.Println(s)                // [1 2 3 4 5 6]
}
