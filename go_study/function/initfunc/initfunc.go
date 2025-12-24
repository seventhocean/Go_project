package main

import "fmt"

var db int

// init函数会在main函数执行前被调用，且可以有多个init函数
func init() {
	db = 10
	fmt.Println("collection db ......success", db)
	fmt.Println("init - 1")
}

func init() {
	fmt.Println("init function called")
}

func init() {
	fmt.Println("init function called")
}

func init() {
	fmt.Println("init function called")
}

func main() {
	fmt.Println("main function called")
}
