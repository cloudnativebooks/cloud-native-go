package _select

func ExampleSelectForChan() {
	c := make(chan string, 1)
	// 由于chan中没有数据，select只会响应写分支
	SelectForChan(c)
	<-c
	// Output:
	// sent Hello
}

func ExampleSelectForChan2() {
	c := make(chan string, 1)
	c <- "Hello"
	// 由于chan中已有1个数据（缓冲区满，select只会响应读分支）
	SelectForChan(c)
	// Output:
	// received Hello
}

func ExampleSelectForChan3() {
	// c := make(chan string)
	// var c chan string

	// 由于chan无缓冲区，且无协程向其写入和读出，那么select永久阻塞
	// SelectForChan(c) // 注释掉，永远阻塞
	// Output:
	// received Hello
}

// 两个case 语句随机执行
func ExampleSelectForChan4() {
	c := make(chan string, 2)
	c <- "Hello"
	SelectForChan(c)
}

// 3个case 语句随机执行
func ExampleSelectAssign() {
	c := make(chan string, 1)
	c <- "Hello"
	SelectAssign(c)
}

// 关闭的channel，三个case语句都有机会执行，但只有返回两个值时才能感知到close。
func ExampleSelectAssign2() {
	c := make(chan string)
	close(c)
	SelectAssign(c)
}

func ExampleSelectDefault() {
	SelectDefault()
	// Output:
	// no data found in default
}

func ExampleSelectChanStillBlock() {
	SelectChanStillBlock()
	// Output:
	// time out
}
