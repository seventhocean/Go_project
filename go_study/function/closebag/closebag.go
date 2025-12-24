package main

import (
	"time"
)

func awaitSomething(awaittime int) func(...int) int {
	time.Sleep(time.Duration(awaittime) * time.Second)
	return func(nums ...int) (sum int) {
		for _, i := range nums {
			sum += i
		}
		return sum
	}
}

func main() {
	// 调用闭包函数
	t1 := time.Now()
	sum := awaitSomething(2)(1, 2, 3)
	t2 := time.Since(t1)
	println(sum, t2)

}
