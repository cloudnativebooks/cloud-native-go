package errors

import (
	"errors"
	"fmt"
	"os"
)

func ExampleCreateBasicError() {
	err := errors.New("this is demo error")

	basicErr := fmt.Errorf("some context: %v", err)           // 使用 %v
	if _, ok := basicErr.(interface{ Unwrap() error }); !ok { // 如果errBasic没有实现Unwrap接口
		fmt.Print("basicErr is a errorString")
	}
	// Output:
	// basicErr is a errorString
}

func ExampleCreateWrapError() {
	err := errors.New("this is demo error")

	wrapError := fmt.Errorf("some context: %w", err)          // 使用%w
	if _, ok := wrapError.(interface{ Unwrap() error }); ok { // 如果wrapError实现了Unwrap接口
		fmt.Print("wrapError is a wrapError")
	}
	// Output:
	// wrapError is a wrapError
}

/*
// 编译错误：Errorf call has more than one error-wrapping directive %w
func ExampleTwoVerb() {
	err := errors.New("this is demo error")

	wrapError := fmt.Errorf("some context: %w, %w", err, err)
	fmt.Printf(wrapError.Error())
	// Output:
	//
}
*/

/*
编译错误： Errorf format %w has arg "some string" of wrong type string
func ExampleNonError() {
	wrapError := fmt.Errorf("some context: %w", "some string")
	fmt.Printf(wrapError.Error())
	// Output:
	//
}
*/

func ExampleUnwrap() {
	err := fmt.Errorf("write file error: %w", os.ErrPermission)
	if errors.Unwrap(err) == os.ErrPermission {
		fmt.Printf("permission denied")
	}
	// Output:
	// permission denied
}

func ExampleUnwrapLoop() {
	err1 := fmt.Errorf("write file error: %w", os.ErrPermission)
	err2 := fmt.Errorf("write file error: %w", err1)

	err := err2
	for {
		if err == os.ErrPermission {
			fmt.Printf("permission denied")
			break
		}
		if err = errors.Unwrap(err); err == nil {
			break
		}
	}
	// Output:
	// permission denied
}

func ExampleIs() {
	err1 := fmt.Errorf("write file error: %w", os.ErrPermission)
	err2 := fmt.Errorf("write file error: %w", err1)

	if errors.Is(err2, os.ErrPermission) {
		fmt.Printf("permission denied")
	}
	// Output:
	// permission denied
}

func ExampleAssertChanErrorWithoutAs() {
	err := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}

	err2 := fmt.Errorf("some context: %w", err)
	if e, ok := err2.(*os.PathError); ok { // 判断失效
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v\n", e.Op, e.Path, e.Err)
	}
	// Output:
	//
}

func ExampleAssertChanWithAs() {
	err := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}

	err2 := fmt.Errorf("some context: %w", err)
	var target *os.PathError
	if errors.As(err2, &target) { // 逐层剥离err2并检测是否是os.PathError类型
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v\n", target.Op, target.Path, target.Err)
	}
	// Output:
	// it's an os.PathError, operation: write, path: /root/demo.txt, msg: permission denied
}
