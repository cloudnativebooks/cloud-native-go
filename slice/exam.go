package _slice

import "fmt"

/*
下面函数输出什么？
*/
func SliceCap() {
	var array [10]int
	var slice = array[5:6]

	fmt.Printf("len(slice) = %d\n", len(slice))
	fmt.Printf("cap(slice) = %d\n", cap(slice))
}

/*
下面函数输出什么？
单选：
- A： [2, 3] [2, 3, 4]
- B： [1, 2] [1, 2, 3]
- C： [1, 2] [2, 3, 4]
- D： [2, 3, 1] [2, 3, 4, 1]
*/
func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}

func SlicePrint() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	SliceRise(s1)
	SliceRise(s2)
	fmt.Println(s1, s2)
}

/*
关于下面函数的描述，正确的是？

单选：
- A： 函数操作nil切片会panic
- B： 编译错误，不可以对切片元素取址
- C： 函数输出"true"
- D： 函数输出"false"
*/
func SliceExtend() {
	var slice []int
	s1 := append(slice, 1, 2, 3)
	s2 := append(s1, 4)

	fmt.Println(&s1[0] == &s2[0])
}

/*
下面函数输出什么？
*/
func SliceExpress() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}
