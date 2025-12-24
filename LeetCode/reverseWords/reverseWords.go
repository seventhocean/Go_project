package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var res []string
	i := len(s) - 1
	for i >= 0 {
		if s[i] == ' ' {
			i--
			continue
		}
		j := i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res = append(res, s[i+1:j+1])

	}
	return strings.Join(res, " ")
}

func main() {
	var result string
	println("Enter a string:")
	fmt.Scanln(&result)
	println(reverseWords(result))
}
