package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello, Go!")
}

func param1(id int) int {
	fmt.Println("id =", id)
	return 1
}
func param2(name string) {
	fmt.Println("name =", name)
}

func param3(name string, id int) {
	fmt.Println("name =", name, ", id =", id)
}

func add(numberList ...int) {
	var sum int
	for _, value := range numberList {
		sum += value
	}
	fmt.Println("sum =", sum)
	fmt.Println("number of parameters =", len(numberList))
	fmt.Printf("numberList = %v\n", numberList)
}

func main() {
	sayHello()
	param1(100)
	param2("gaoyuan")
	param3("gaoyuan", 200)
	add(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3)
}
