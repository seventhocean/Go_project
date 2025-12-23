// map
package main

func majorityElement(nums []int) int {
	n := len(nums)
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] > n/2 {
			return num
		}
	}
	return -1
}

// Boyer-Moore 投票算法

func majorityElement1(nums []int) int {
	var candidate int
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
