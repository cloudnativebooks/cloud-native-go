package errors

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/*
官方博客中出现的一个笔误：
原文链接：https://blog.golang.org/go1.13-errors
原文如下：
type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

if e, ok := err.(*NotFoundError); ok {
    // e.Name wasn't found
}

其中“e.Name wasn't found” 错误，实际应该是“e.Name was found”

修复该笔误的PR：https://github.com/golang/blog/pull/33
更新：是笔者理解错了，`e.Name wasn't found`并不是指err中没有e.Name成员（从而不是NotFoundError），而是指e.Name的值没有找到。
*/

func ExampleNotFoundError_Error() {
	err1 := errors.New("this is not NotFoundError")
	if e, ok := err1.(*NotFoundError); ok {
		fmt.Printf("err1 is a NotFoundError, e.Name=%s", e.Name)
	}

	err2 := NewNotFoundError("err2")

	if e, ok := err2.(*NotFoundError); ok {
		fmt.Printf("err2 is a NotFoundError, e.Name=%s", e.Name)
	}

	// OutPut:
	// err2 is a NotFoundError, e.Name=err2
}

// errors.New() 性能测试
// [root@ecs-d8b6 errors]# benchstat new.txt
// name                time/op
// MakeByErrorsNew-12  0.24ns ± 1%
func BenchmarkMakeByErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeByErrorsNew()
	}
}

// fmt.Errorf() 性能测试
// [root@ecs-d8b6 errors]# benchstat old.txt
// name                time/op
// MakeByFmtErrorf-12  80.9ns ± 1%
func BenchmarkMakeByFmtErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MakeByFmtErrorf()
	}
}

// go test ./errors -run=ExampleAssertError
func ExampleAssertError() {
	err1 := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}
	AssertError(err1)

	err2 := fmt.Errorf("not an os.PathError")
	AssertError(err2)

	// Output:
	// it's an os.PathError, operation: write, path: /root/demo.txt, msg: permission denied
	// permission denied
}

// go test ./errors -run=ExampleWriteFile
func ExampleWriteFile() {
	err := WriteFile("a.txt")
	if err == os.ErrPermission {
		fmt.Printf("permission denied")
	}
	// Output:
	//
}
