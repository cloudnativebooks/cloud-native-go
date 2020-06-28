package recover

import "fmt"

// 考察点：recover之后，流程不会回到panic后面继续执行
func RecoverDemo1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic("demo")
	fmt.Println("B")
}

// 考察点：剩余defer中仍可以执行
func RecoverDemo2() {
	defer func() {
		fmt.Println("C")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic("demo")
	fmt.Println("B")
}

// 考察点：非顶层函数无法处理defer
func RecoverDemo3() {
	defer func() {
		func() {
			if err := recover(); err != nil {
				fmt.Println("A")
			}
		}()
	}()

	panic("demo")
	fmt.Println("B")
}

// 考察点：两次recover 第二次无效
func RecoverDemo4() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("B")
		}
	}()

	panic("demo")
	fmt.Println("C")
}

// 考察点：函数返回零值，且感知不到panic
func RecoverDemo5() {
	foo := func() int {
		defer func() {
			recover()
		}()

		panic("demo")

		return 10
	}

	ret := foo()
	fmt.Println(ret)
}

// 考察点：panic(nil)时，recover()返回nil
func RecoverDemo6() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic(nil)
	fmt.Println("B")
}
