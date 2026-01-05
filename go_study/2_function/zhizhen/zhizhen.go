package main

import "fmt"

func add(num int) int {
	fmt.Println(&num)
	num += 10
	return num
}

func SetValue(num *string) {
	*num = "jinjin"
	fmt.Println("SetValue:", *num)
}
func main() {
	// 指针作为函数参数传递, 修改指针指向的值, 而不是指针本身
	num := 20
	fmt.Println(&num)
	result := add(num)
	fmt.Println(result)
	fmt.Println(num)

	// 引用类型作为函数参数传递，修改引用类型指向的值，而不是引用类型本身
	var str string = "gaoyuan"
	SetValue(&str)
	fmt.Println("TRUE:", str)

}
