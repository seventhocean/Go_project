package main

import "fmt"

func main() {

	// defer语句会将函数推迟到外层函数返回之后执行,
	// 即使外层函数发生panic，defer语句仍然会被执行
	fmt.Println("main func start")
	defer fmt.Println("defer func 1")
	defer func() {
		fmt.Println("defer anonymous func")
	}()
	defer fmt.Println("defer func 2")
	defer fmt.Println("defer func 3")
	// 输出顺序与defer语句的书写顺序相反，后写的先执行
	fmt.Println("main func end")
}
