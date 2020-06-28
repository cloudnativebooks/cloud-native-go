package _defer

import "fmt"

// defer执行顺序后进先出
func ExampleDeferOrder() {
	DeferDemo1()
	// Output:
	// 43210
}

func ExampleDeferPinVar() {
	DeferDemo2()
	// Output:
	// 1
}

func ExampleDeferPinFun() {
	DeferDemo3()
	// Output:
	// 1
}

func ExampleDeferPinAddr() {
	DeferDemo4()
	// Output:
	// 1023
}

func ExampleDeferReturn() {
	result := DeferDemo5()
	fmt.Print(result)
	// Output:
	// 2
}

func ExampleDeferNest() {
	DeferDemo6()
	// Output:
	// AB
}
