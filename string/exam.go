package _string

import "fmt"

/*
关于下面函数中字符串长度描述正确的是？
单选：
- A：字符串长度表示字符个数，长度为2
- B：字符串长度表示Unicode编码字节数，长度大于2
- C：不可以针对中文字符计算长度
- D：不确定，与运行环境有关
*/
func StringExam1() {
	var s string
	s = "中国"
	fmt.Println(len(s))
}

/*
关于字符串的描述正确的是？

单选：
- A：如果string为nil，那么其长度为0
- B：不存在值为nil的string
*/

/*
下面关于字符串的描述正确的是？
单选：
- A： 可以使用下标修改string的内容
- B： 字符串不可以修改
*/
