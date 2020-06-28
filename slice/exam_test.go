package _slice

func ExampleSliceCap() {
	SliceCap()
	// Output:
	// len(slice) = 1
	// cap(slice) = 5
}

func ExampleSlicePrint() {
	SlicePrint()
	// Output:
	// [1 2] [2 3 4]
}

func ExampleSliceExtend() {
	SliceExtend()
	// Output:
	// true
}

func ExampleSliceExpress() {
	SliceExpress()
	// Output:
	// len(pollorder) =  5
	// cap(pollorder) =  5
	// len(lockorder) =  5
	// cap(lockorder) =  5
}
