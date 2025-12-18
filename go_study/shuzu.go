package main

import (
	"fmt"
)

func main() {

	var namelist [3]string = [3]string{"gaoyuan", "wangshan", "jinjin"}
	fmt.Println(namelist)
	fmt.Println(namelist[1])
	fmt.Println(namelist[len(namelist)-1])
}
