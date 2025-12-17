package main

import (
	"fmt"
)

func main() {
	fmt.Println("output.go")
	fmt.Printf("%d", 12)
	fmt.Printf("%.2f\n", 3.1415926)
	fmt.Printf("%T %T\n", "nihao", 3.2)
	fmt.Printf("%v\n", "suibian")
	fmt.Printf("%v\n", "")
	fmt.Printf("%#v\n", "")

	var f = fmt.Sprintf("%d + %d = %d", 1, 2, 3)
	fmt.Printf("%s", f)
}
