package _range

import "testing"

func BenchmarkFindMonkey(b *testing.B) {
	s := []string{"apple", "peach", "lion", "elephant", "monkey"}

	for i := 0; i < b.N; i++ {
		FindMonkey(s)
	}
}

func BenchmarkFindMonkeyImproved(b *testing.B) {
	s := []string{"apple", "peach", "lion", "elephant", "monkey"}

	for i := 0; i < b.N; i++ {
		FindMonkeyImproved(s)
	}
}

func ExamplePrintSlice() {
	PrintSlice()
	// Output:
	// 3
	// 3
	// 3
}

func ExampleRangeTimer() {
	// RangeTimer() // 会阻塞到go test 超时
	// Output:
}

func ExampleRangeDemo() {
	RangeDemo()
	// Output:
}
