package _iota

import "fmt"

func Exam1() {
	const (
		mutexLocked = 1 << iota // mutex is locked
		mutexWoken
		mutexStarving
		mutexWaiterShift      = iota
		starvationThresholdNs = 1e6
	)

	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	fmt.Println(starvationThresholdNs)
}

const (
	// 常量的注释（文档）
	a, b = iota, iota // 常量的行注释
)
