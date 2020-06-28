package errors

import (
	"errors"
	"fmt"
	"reflect"
)

// 内容为空的error也是异常。
func ExampleEmptyError() {
	err := errors.New("")
	if err != nil {
		fmt.Printf("empty error still is an error")
	}
	// OutPut:
	// empty error still is an error
}

func ExampleFormatVerb() {
	err := errors.New("not found")
	err1 := fmt.Errorf("some context: %v", err)
	err2 := fmt.Errorf("some context: %w", err)

	if err1 == err2 {
		fmt.Printf("two errors are equal\n")
	} else {
		fmt.Printf("two errors are different\n")
	}

	if err1.Error() != err2.Error() {
		panic("two errors's content should be same")
	}

	fmt.Printf("%s\n", reflect.TypeOf(err1).String())
	fmt.Printf("%s\n", reflect.TypeOf(err2).String())

	// Output:
	// two errors are different
	// *errors.errorString
	// *fmt.wrapError
}

func ExampleUnwrapBasicError() {
	err := errors.New("not found")

	err = errors.Unwrap(err)
	switch err {
	case nil:
		fmt.Printf("err is nil")
	default:
		fmt.Printf("err is non-nil")
	}

	// Output:
	// err is nil
}
