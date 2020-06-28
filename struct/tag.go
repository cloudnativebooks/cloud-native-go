package _struct

import (
	"fmt"
	"reflect"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}

func PrintTag() {
	t := TypeMeta{}
	ty := reflect.TypeOf(t)

	for i := 0; i < ty.NumField(); i++ {
		fmt.Printf("Field: %s, Tag: %s\n", ty.Field(i).Name, ty.Field(i).Tag.Get("json"))
	}
}
