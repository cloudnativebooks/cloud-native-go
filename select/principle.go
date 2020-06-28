package _select

import (
	"fmt"
	"unsafe"
)

type scase struct {
	c           string         // chan
	elem        unsafe.Pointer // data element
	kind        uint16
	pc          uintptr // race pc (for race detector / msan)
	releasetime int64
}

func SelectGo(cas0 *scase, order0 *uint16, ncases int) {
	cas1 := (*[1 << 16]scase)(unsafe.Pointer(cas0))
	fmt.Printf("len(cas1)=%d, cap(cas1)=%d\n", len(cas1), cap(cas1))

	order1 := (*[1 << 17]uint16)(unsafe.Pointer(order0))
	fmt.Printf("len(order1)=%d, cap(order1)=%d\n", len(order1), cap(order1))

	scases := cas1[:ncases:ncases]
	fmt.Printf("len(scases)=%d, cap(scases)=%d\n", len(scases), cap(scases))
	pollorder := order1[:ncases:ncases]
	fmt.Printf("len(pollorder)=%d, cap(pollorder)=%d\n", len(pollorder), cap(pollorder))
	lockorder := order1[ncases:][:ncases:ncases]
	fmt.Printf("len(lockorder)=%d, cap(lockorder)=%d\n", len(lockorder), cap(lockorder))
}

func WriteNilChannle() {
	var c chan string
	select {
	case c <- "hello": // 向值为nil的管道写数据不会触发panic
		fmt.Println("true")
	default:
		fmt.Println("default")
	}
}
