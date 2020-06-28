package sugar

import (
	"fmt"
)

// rule := "Short variable declarations"  // syntax error: non-declaration statement outside function body

func fun1() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}

func fun2(i int) {
	// i := 0  // 无法通过编译：no new variable on left side of :=
	fmt.Println(i)
}

func fun3() {
	i, j := 0, 0
	if true {
		j, k := 1, 1
		fmt.Printf("j = %d, k = %d\n", j, k)
	}
	fmt.Printf("i = %d, j = %d\n", i, j)
}

func nextField() (field int, err error) {
	return 1, nil
}

/*
下面代码变量未使用编译不通过，仅用于说明用法
func Redeclare() {
    field, err:= nextField()   // 1号err

    if field == 1{
        field, err:= nextField()     //　2号err
        newField, err := nextField() //  3号err
        ...
    }
    ...
}
*/
func SugerPackage() {
	fun1()
	fun3()
}
