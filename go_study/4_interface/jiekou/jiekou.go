package main

import (
	"fmt"
)

// 定义一个接口 Bird，包含 Sing 和 Walk 方法，Chicken 类型实现了该接口，因此可以赋值给 Bird 接口变量，从而实现多态。
type Bird interface {
	Sing()
	Walk()
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Chicken) Walk() {
	fmt.Println(c.Name, "在走路")
}

type Fish struct {
	Name string
}

func (f Fish) Sing() {
	fmt.Println(f.Name, "在唱歌")
}
func (f Fish) Walk() {
	fmt.Println(f.Name, "在走路")
}

func SingandWalk(b Bird) {
	switch server := b.(type) {
	case Chicken:
		fmt.Println(server.Name)
	case Fish:
		fmt.Println(server.Name)
	default:
		fmt.Println("未知类型")
	}

	b.Sing()
	b.Walk()
}

func main() {
	chicken := Chicken{Name: "ikun"}
	fish := Fish{Name: "小黄鱼"}
	SingandWalk(chicken)
	SingandWalk(fish)

}
