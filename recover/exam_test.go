package recover

func ExampleRecoverDemo1() {
	RecoverDemo1()
	// Output:
	// A
}

func ExampleRecoverDemo2() {
	RecoverDemo2()
	// Output:
	// A
	// C
}

func ExampleRecoverDemo3() {
	defer func() {
		recover()
	}()
	RecoverDemo3()
	// Output:
	//
}

func ExampleRecoverDemo4() {
	RecoverDemo4()
	// Output:
	// B
}

func ExampleRecoverDemo5() {
	RecoverDemo5()
	// Output:
	// 0
}

func ExampleRecoverDemo6() {
	RecoverDemo6()
	// Output:
	//
}
