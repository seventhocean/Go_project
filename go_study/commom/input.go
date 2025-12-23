package main

import (
	"fmt"
)

func main() {

	fmt.Printf("please input name?\n")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)

	fmt.Printf("please input age?\n")
	var age int
	// fmt.Scan(&age)
	// fmt.Println(age)
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}
