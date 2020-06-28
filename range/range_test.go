package _range

func ExampleRangeArray() {
	RangeArray()
	// Output:
	// index: 0 value:1
	// index: 1 value:2
	// index: 2 value:3
}

func ExampleRangeArrayPointer() {
	RangeArrayPointer()
	// Output:
	// index: 0 value:1
	// index: 1 value:2
	// index: 2 value:3
}

func ExampleRangeSlice() {
	RangeSlice()
	// Output:
	// index: 0 value:1
	// index: 1 value:2
	// index: 2 value:3
}

// TODO(RainbowMango): 如果字符串中包含空格，那么Example测试将失败，怀疑是Bug
func ExampleRangeString() {
	RangeString()
	// Output:
	// index: 0, value: H
	// index: 1, value: e
	// index: 2, value: l
	// index: 3, value: l
	// index: 4, value: o
}

func ExampleRangeStringUniCode() {
	RangeStringUniCode()
	// Output:
	// index: 0, value: 中
	// index: 3, value: 国
}

func ExampleRangeMap() {
	RangeMap()
	// Unordered output:
	// key: animal, value: monkey
	// key: fruit, value: apple
}

func ExampleRangeChannel() {
	RangeChannel()
	// Output:
	// element: Hello
	// element: World

}
