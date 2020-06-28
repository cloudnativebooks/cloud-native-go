package _map

import (
	"fmt"
)

// MapInitByLiteral 演示使用字面量初始化map
func MapInitByLiteral() {
	m := map[string]int{
		"apple":  2,
		"banana": 3,
	}

	for k, v := range m {
		fmt.Printf("%s-%d\n", k, v)
	}
}

// MapInitByMake 演示使用make初始化map
func MapInitByMake() {
	m := make(map[string]int, 10) // 或 m := make(map[string]int)
	m["apple"] = 2
	m["banana"] = 3

	for k, v := range m {
		fmt.Printf("%s-%d\n", k, v)
	}
}

// 演示增删改查
func MapCRUD() {
	m := make(map[string]string, 10)
	m["apple"] = "red"     // 添加
	m["apple"] = "green"   // 修改
	delete(m, "apple")     // 删除
	v, exist := m["apple"] // 查询
	if exist {
		fmt.Printf("apple-%s\n", v)
	}
}

func EmptyMap() {
	var m1 map[string]int
	m2 := make(map[string]int)

	fmt.Printf("len(m1) = %d\n", len(m1)) // 0
	fmt.Printf("len(m2) = %d\n", len(m2)) // 0

	// nil map 可以查询
	v, exist := m1["apple"]
	if exist {
		fmt.Println(v) // 永不可达
	}
}
