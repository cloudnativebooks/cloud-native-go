package _string

import (
	"fmt"
)

// StringFormat 演示`%s`和`%q`格式的区别
// 更多参考信息：https://github.com/kubernetes-sigs/kubefed/pull/1216
func StringFormat() {
	var s string
	s = "https://example.com:port80"
	fmt.Printf("format-s: %s\n", s)
	fmt.Printf("format-q: %q\n", s)
}
