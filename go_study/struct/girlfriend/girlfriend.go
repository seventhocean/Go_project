package main

import "fmt"

type Class struct {
	Name   string
	Room   string
	Size   int
	Number int
}
type Student struct {
	// 嵌套结构体,表示Student结构体包含了Class结构体的所有字段,相当于继承了Class结构体
	Class
	Name string
}

// 给Student结构体绑定一个Study方法,s是Student类型的接收者,study是方法名,==> func (接收者 结构体类型) 方法名() {}
func (s Student) Study() {
	fmt.Printf("%s 正在学习\n", s.Name)
}

func (s Student) Info() {
	// 访问嵌套结构体的字段,直接使用字段名即可,不需要通过嵌套结构体名访问,如s.Class.Room
	fmt.Printf("学生姓名:%s,教室:%s,教室大小:%d,教室编号:%d\n", s.Name, s.Room, s.Size, s.Number)
}

//值类型接收者,无法修改嵌套结构体的字段值
func (s Student) SetClassName(name string) {
	s.Class.Name = name
	fmt.Printf("Class in func %s\n", s.Class.Name)
	fmt.Printf("Class in func %p\n", &s.Class.Name)

}

func (s *Student) SetClassRoom(room string) {
	s.Class.Room = room
	fmt.Printf("Class in func %s\n", s.Class.Room)
	fmt.Printf("Class in func %p\n", &s.Class.Room)
}

func main() {
	// 创建Class结构体实例,并初始化字段,赋值给变量c1
	c1 := Class{
		Name:   "三年一班",
		Size:   40,
		Room:   "201",
		Number: 1,
	}
	// 创建Student结构体实例,并初始化字段,赋值给变量s1,嵌套结构体使用变量c1,表示s1的Class字段是c1,即s1继承了c1的所有字段
	s1 := Student{Name: "gaoyuan", Class: c1}
	s1.Study()
	s1.Info()
	// 也可以直接在创建Student结构体实例时,初始化嵌套结构体的字段,不需要先创建Class结构体实例,如下面的s2
	s2 := Student{Name: "jinjin", Class: Class{Room: "102", Size: 60, Number: 2}}
	s2.Study()
	s2.Info()
	fmt.Println("-------------------------------------")
	s3 := Student{Class: Class{Name: "mingming"}}
	// 修改嵌套结构体的字段值，调用SetClassName方法，但由于SetClassName方法的接收者是值类型，无法修改嵌套结构体的字段值，所以修改无效
	fmt.Println("修改班级名称前:")
	fmt.Printf("班级名称: %s\n", s3.Class.Name)
	fmt.Printf("班级名称: %p\n", &s3.Class.Name)
	fmt.Println("修改班级名称后:")
	s3.SetClassName("三年一班")
	fmt.Printf("班级名称: %s\n", s3.Class.Name)
	fmt.Printf("班级名称: %p\n", &s3.Class.Name)
	fmt.Println("-------------------------------------")
	s4 := Student{Class: Class{Room: "303"}}
	// 修改嵌套结构体的字段值，调用SetClassRoom方法，由于SetClassRoom方法的接收者是指针类型，可以修改嵌套结构体的字段值，所以修改有效
	fmt.Println("修改教室前:")
	fmt.Printf("教室: %s\n", s4.Class.Room)
	fmt.Printf("教室: %p\n", &s4.Class.Room)
	fmt.Println("修改教室后:")
	s4.SetClassRoom("404")
	fmt.Printf("教室: %s\n", s4.Class.Room)
	fmt.Printf("教室: %p\n", &s4.Class.Room)
}
