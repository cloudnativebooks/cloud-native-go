package _chan

import "fmt"

func ChanInitDeclare() {
	var ch chan int // 声明管道
	fmt.Println(ch == nil)
}

func ChanInitMake() {
	ch1 := make(chan string)    // 无缓冲管道
	ch2 := make(chan string, 5) // 带缓冲管道

	fmt.Println(ch1 == nil)
	fmt.Println(ch2 == nil)
}

func ChanOperator() {
	ch := make(chan int, 10)
	ch <- 1   // 数据流入管道
	d := <-ch // 数据流出管道
	fmt.Println(d)
}

func ChanParamRW(ch chan int) {
	// 管道可读写
}

func ChanParamR(ch <-chan int) {
	// 只能从管道读取数据
}

func ChanParamW(ch chan<- int) {
	// 只能向管道写入数据
}

// ReadChanReturnValue 用于测试读取管道时的第二个参数到底是什么含义？
// 第二个参数代表是否成功读取了数据。（不代表是否关闭）
func ReadChanReturnValue() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	fmt.Printf("length of channel is: %d\n", len(ch))

	_, ok := <-ch
	fmt.Printf("second return value before channel closed is: %v\n", ok) // true

	close(ch)

	_, ok = <-ch
	fmt.Printf("second return value after channel(have data) closed is: %v\n", ok) // true

	_, ok = <-ch
	fmt.Printf("second return value after channel(no data) closed is: %v\n", ok) // false
}
