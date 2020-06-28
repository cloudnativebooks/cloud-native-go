package recover

import "fmt"

func ExampleNamedRecover() {
	NamedRecover()
	// Output:
	// catch a recover from named-defer-function
}

func ExampleGoExit() {
	// GoExit() // go test ./recover 执行时会让hang住，暂不确定是否Go bug.
	// Output:
	// runtime.Goexit will call defer like panic.
}

func ExampleGoExitRecover() {
	fmt.Printf("program will hanging there!")
	// GoExitRecover() // go test ./recover 执行时会让hang住，暂不确定是否Go bug.
	// Output:
	//
}
