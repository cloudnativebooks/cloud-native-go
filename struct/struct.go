package _struct

import (
	"fmt"
	"reflect"
)

// 嵌入字段类型名将作为字段名存在
func EmbeddedField() {
	type A struct {
		a1 int
	}
	type B struct {
		b1 int
		A  // embedded field
	}

	b := B{}
	t := reflect.TypeOf(b)
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field-%d, Name: %s, isEmbedded: %v\n", i, t.Field(i).Name, t.Field(i).Anonymous)
	}
}

// 非空fields名称必须唯一
func FieldsUnique() {
	type A struct {
		a1 int
		_  int // blank field
		_  int // blank field
	}
}
