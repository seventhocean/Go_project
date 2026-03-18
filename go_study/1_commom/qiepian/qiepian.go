package main

import (
	"fmt"
	"sort"
)

func main() {
	var list []string

	list = append(list, "gaoyuan")
	list = append(list, "yangjian")
	list = append(list, "jinjin")
	fmt.Println(list)
	fmt.Println(list[1])
	fmt.Println(list[len(list)-1])
	fmt.Println(list == nil)

	var namelist = []string{}
	fmt.Println(namelist)

	namelist1 := make([]string, 0)
	fmt.Println(namelist1)

	array := [4]int{1, 2, 3, 4}
	slices := array[:]
	fmt.Println(slices)

	slices1 := array[1:3]
	fmt.Println(slices1)

	var ints = []int{2, 5, 4, 8, 1, 5, 2, 4, 6}
	sort.Ints(ints)
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)

}
