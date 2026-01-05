package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"` // 指定JSON序列化时的字段名称
	Age      int    `json:"age"`
	Password string `json:"-"` // 忽略该字段
}

// struct标签：
// 1. 可以通过反射获取struct标签的信息，从而实现一些特殊的功能，比如ORM框架中通过struct标签来映射数据库字段。
// 2. 可以通过struct标签来控制序列化和反序列化的行为，比如在JSON序列化时忽略某些字段，或者指定字段的名称。
// 3. 可以通过struct标签来实现数据验证，比如指定字段的最大长度、最小值等。
func main() {
	user := User{Name: "少夫人", Age: 23, Password: "123456"}
	byteDate, _ := json.Marshal(user)
	fmt.Println(string(byteDate))
}
