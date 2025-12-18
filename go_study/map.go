package main

import (
	"fmt"
)

func main() {
	var usermap = map[int]string{
		1: "gaoyuan",
		2: "wangshan",
		3: "jinjin",
	}
	fmt.Println(usermap)
	fmt.Println(usermap[2])
	fmt.Printf("%#v\n", usermap[3])
	fmt.Printf("%v\n", usermap[3])

	usermap[1] = "yangyuan"
	fmt.Println(usermap)
	delete(usermap, 2)
	fmt.Println(usermap)

}
