package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
	}
	if age > 0 && age <= 18 {
		fmt.Println("未成年")
	}
	if age > 18 && age <= 35 {
		fmt.Println("青年")
	}
	if age > 35 {
		fmt.Println("中年")
	}
}
