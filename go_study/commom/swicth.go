package main

import "fmt"

func main() {
	fmt.Println("请输入星期数字：")
	var week int
	fmt.Scan(&week)

	switch week {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	case 4:
		fmt.Println("周四")
	case 5:
		fmt.Println("周五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("错误")
	}
}
