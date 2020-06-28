package _slice

// 测试拷贝切片时遍历和内置函数性能差异
func CopyByIteration(origin []string) []string {
	ret := make([]string, len(origin))
	for i := range origin {
		ret[i] = origin[i]
	}

	return ret
}

// 测试拷贝切片时遍历和内置函数性能差异
func CopyByBuiltIn(origin []string) []string {
	ret := make([]string, len(origin))
	copy(ret, origin)
	return ret
}
