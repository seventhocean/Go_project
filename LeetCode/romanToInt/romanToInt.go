package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if i < n-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			total -= romanMap[s[i]]
		} else {
			total += romanMap[s[i]]
		}
	}
	return total
}

// Example usage:

func main() {
	var result string
	fmt.Print("Enter a Roman numeral: ")
	fmt.Scanln(&result)
	fmt.Println(romanToInt(result))

}
