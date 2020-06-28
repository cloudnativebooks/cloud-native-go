package sugar

// 可变参数，调用时可以不填值
func ExampleGreetingWithoutParameter() {
	Greeting("nobody")
	// OutPut:
	// Nobody to say hi.
}

// 可变参数传值
func ExampleGreetingWithParameter() {
	Greeting("hello:", "Joe", "Anna", "Eileen")
	// OutPut:
	// hello: Joe
	// hello: Anna
	// hello: Eileen
}

// 可变参数传切片
func ExampleGreetingWithSlice() {
	guest := []string{"Joe", "Anna", "Eileen"}
	Greeting("hello:", guest...)
	// OutPut:
	// hello: Joe
	// hello: Anna
	// hello: Eileen
}
