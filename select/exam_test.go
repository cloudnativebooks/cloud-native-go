package _select

// 结果：输出随机
func ExampleSelectExam1() {
	SelectExam1()
}

// 结果：函数阻塞
func ExampleSelectExam2() {
	SelectExam2()
}

func ExampleSelectExam3() {
	SelectExam3()
	// Output:
	// 1
}

func ExampleSelectExam4() {
	SelectExam4()
	// Output:
	// 1
}

// 结果：函数阻塞
func ExampleSelectExam5() {
	SelectExam5()
}

func ExampleSelectExam6() {
	SelectExam6()
	// Output:
	// default
}
