package main

import (
	"fmt"

	"go_study/1_commom/version"
)

func hello() {
	fmt.Println(version.Version)
	fmt.Println(version.Gaoyuan)

}

func main() {
	var name string = "gaoyuan"
	fmt.Println(name)

	var name1 string = "gaoyuan1"
	var name2 = "gaoyuan2"
	name3 := "gaoyuan3"
	fmt.Printf("%s %s %s", name1, name2, name3)

	var a1, a2 = 1, 2
	fmt.Println(a1, a2)
	hello()

}
