package panic

import "fmt"

func foo() {
	defer fmt.Print("A")
	defer fmt.Print("B")

	fmt.Print("C")
	panic("demo")
	defer fmt.Print("D")
}

// 考查点：panic会触发单个函数中的所有defer函数
func PanicDemo1() {
	defer func() {
		recover()
	}()

	foo()
}

// 考查点：panic会触发整个协程中的所有defer函数
func PanicDemo2() {
	defer func() {
		recover()
	}()

	defer func() {
		fmt.Print("1")
	}()

	foo()
}

// 考察点： panic不会触发其他协程的defer函数
func PanicDemo3() {
	defer func() {
		fmt.Print("demo")
	}()

	go foo()
}

func PanicDemo4() {
	defer func() {
		recover()
	}()

	defer fmt.Print("A")

	defer func() {
		fmt.Print("B")
		panic("panic in defer")
		fmt.Print("C")
	}()

	panic("panic")

	fmt.Print("D")
}
