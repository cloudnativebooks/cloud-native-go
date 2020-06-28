package panic

import "fmt"

func FooPanicProcess() {
	ret := fooPanic()

	fmt.Printf("continue after a panic-recovered function.\n")
	fmt.Printf("the panic-recovered function return value is zero value. ret = %d\n", ret)
}

func fooPanic() int {
	defer func() {
		// rule: recover之后继续处理当前协程的defer
		fmt.Printf("after recover, continue call defer\n")
	}()
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("panic recovered: %v\n", msg)
		}
	}()

	panic("trigger panic")

	// rule: panic recover之后不会再返回panic点继续执行
	fmt.Printf("after recover, the process can not go back here anymore.")

	return 10
}
