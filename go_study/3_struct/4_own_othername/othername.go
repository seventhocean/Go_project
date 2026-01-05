package main

import "fmt"

// 自定义类型和类型别名的区别，以整数类型为例，定义一个自定义类型 MyCode 和一个类型别名 MyAliasCode，并尝试为它们定义方法。
// 自定义类型可以定义方法，而类型别名不能定义方法，因为类型别名本质上还是原始类型。
// 自定义类型适用于需要为类型添加方法或实现接口的场景，而类型别名适用于简化类型名称或提高代码可读性的场景。
type MyCode int

// 别名类型本质上还是原始类型
type MyAliasCode = int

const (
	SuccessCode  MyCode      = 1
	SuccessAlias MyAliasCode = 1
)

// GetMsg 方法用于将 MyCode 类型转换为对应的字符串信息
func (m MyCode) GetMsg(n int) string {
	switch MyCode(n) {
	case SuccessCode:
		return "成功"
	default:
		return "未知状态"
	}
}

// //类型别名不能定义方法，因为它本质上还是原始类型
// func (n MyAliasCode) GetMsg() string {

// }

func main() {
	fmt.Printf("自定义类型 MyCode: %v, %T\n", SuccessCode, SuccessCode)
	fmt.Printf("类型别名 MyAliasCode: %v, %T\n", SuccessAlias, SuccessAlias)
	fmt.Println("MyCode GetMsg:", SuccessCode.GetMsg(1))
	// fmt.Println("MyAliasCode GetMsg:", SuccessAlias.GetMsg()) // 编译错误

}
