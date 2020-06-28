package _defer

import "fmt"

func DeferDemo1() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func DeferDemo2() {
	var aInt = 1

	defer fmt.Println(aInt)

	aInt = 2
	return
}

func DeferDemo3() {
	var i = 0
	defer func() {
		fmt.Println(i)
	}()
	i++
}

func DeferDemo4() {
	var aArray = [3]int{1, 2, 3}
	defer func(array *[3]int) {
		for i := range array {
			fmt.Print(array[i])
		}
	}(&aArray)
	aArray[0] = 10
}

func DeferDemo5() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func DeferDemo6() {
	defer func() {
		defer func() {
			fmt.Print("B")
		}()
		fmt.Print("A")
	}()
}
