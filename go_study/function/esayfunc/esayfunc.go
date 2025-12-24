package main

import (
	"errors"
	"fmt"
)

// 简单函数
func sayHello() {
	fmt.Println("Hello, Go!")
}

// 带参数函数
func param1(id int) int {
	fmt.Println("id =", id)
	return 1
}

// 带多个参数函数
func param2(name string) {
	fmt.Println("name =", name)
}

// 带多个不同类型参数函数
func param3(name string, id int) {
	fmt.Println("name =", name, ", id =", id)
}

// 可变参数函数
func add(numberList ...int) {
	var sum int
	for _, value := range numberList {
		sum += value
	}
	fmt.Println("sum =", sum)
	fmt.Println("number of parameters =", len(numberList))
	fmt.Printf("numberList = %v\n", numberList)
}

// 无返回值函数
func fun1() {
}

// 有返回值函数
func fun2() int {
	return 1
}

// 多返回值函数
func fun3() (int, error) {
	return 0, errors.New("error")
}

// 命名返回值函数
func fun4() (result string) {
	result = "nihao"
	return result
}

func main() {
	sayHello()
	param1(100)
	param2("gaoyuan")
	param3("gaoyuan", 200)
	// 调用可变参数函数
	add(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3)

	// 调用无返回值函数
	fun1()
	// 调用有返回值函数
	fmt.Println("fun2 =", fun2())
	// 调用多返回值函数
	r, err := fun3()
	fmt.Println("fun3 =", r, err)
	// 调用命名返回值函数
	fmt.Println("fun4 =", fun4())

	var getName = func(name string) string {
		return name
	}
	fmt.Println("getName =", getName("gaoyuan"))

	var setName = func(name string) {
		fmt.Println("name =", name)
	}
	setName("gaoyuan")

}
