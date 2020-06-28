package _struct

import (
	"encoding/json"
	"fmt"
)

/*
关于下面类型的描述，正确的是？
单选：
- A：age为私有字段，只能通过类型方法访问
- B：Name和age字段在可见性上没有区别
- C：age为私有字段，变量初始化时不能直接赋值
- D：当前package中，Name和age字段在可见性上没有区别
*/
type People struct {
	Name string
	age  int
}

func (p *People) SetName(name string) {
	p.Name = name
}

func (p *People) SetAge(age int) {
	p.age = age
}

func (p *People) GetAge() int {
	return p.age
}

/*
关于下面代码的描述，正确的是？
单选：
- A： 编译错误，类型和类型指针不能同时作为方法接收者
- B： SetName()无法修改名字
- C： SetAge()无法修改年龄
- D： SetName()和 SetAge()工作正常
*/
type Kid struct {
	Name string
	Age  int
}

func (k Kid) SetName(name string) {
	k.Name = name
}

func (k *Kid) SetAge(age int) {
	k.Age = age
}

func KidFoo() {
	var k Kid
	k.SetName("Marco")
	k.SetAge(3)
	fmt.Printf("Name: %s, Age: %d\n", k.Name, k.Age)
}

/*
使用标准库json包操作下面结构体时，下面描述正确的是：
```go
f := Fruit{Name: "apple", Color: "red"}
s := `{"Name":"banana","Weight":100}`
```

单选：
- A： json.Marshal(f)时会忽略Color字段
- B： json.Marshal(f)时不会被忽略Color字段
- C： json.Unmarshal([]byte(s), &f)时会忽略Color字段
- D： json.Unmarshal([]byte(s), &f)时会出错，因为Fruit类型没有Weight字段
*/

type Fruit struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

var f = Fruit{Name: "apple", Color: "red"}
var s = `{"Name":"banana","Weight":100}`

func MetaFoo() {
	if b, err := json.Marshal(f); err != nil {
		fmt.Println("unexpected error: ", err.Error())
	} else {
		fmt.Printf("%s\n", string(b))
	}

	fruit := Fruit{}
	s := `{"Name":"banana","Weight":100}`
	if err := json.Unmarshal([]byte(s), &fruit); err != nil {
		fmt.Println("unexpected error: ", err.Error())
	} else {
		fmt.Printf("Name: %s, Color: %s\n", fruit.Name, fruit.Color)
	}
}
