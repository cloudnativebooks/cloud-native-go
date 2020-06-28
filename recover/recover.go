package recover

import (
	"fmt"
	"runtime"
)

// 测试recover能否在具名函数中使用
func ForRecover() {
	if p := recover(); p != nil {
		fmt.Printf("catch a recover from named-defer-function\n")
	}
}

func NamedRecover() {
	defer ForRecover()

	panic("this is a panic")
}

// 测试runtime.Goexit()会不会触发 defer
func GoExit() {
	defer func() {
		fmt.Printf("runtime.Goexit will call defer like panic.\n")
	}()
	runtime.Goexit()
}

func GoExitRecover() {
	defer func() {
		if err := recover(); err == nil {
			fmt.Printf("runtime.Goexit() can not be recovered.\n")
		}
	}()
	runtime.Goexit()
}
