package rotate

// ---使用额外数组的方法---

// func rotate(nums []int, k int) {
// 	n := len(nums)
// 	if n == 0 {
// 		return
// 	}
// 	k = k % n // 处理 k >= n 的情况

// 	// 创建新数组
// 	result := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		result[(i+k)%n] = nums[i]
// 	}

// 	// 复制回原数组
// 	copy(nums, result)
// }

// ---另一种方法：使用反转---
// func rotate(nums []int, k int) {
// 	n := len(nums)
// 	if n == 0 {
// 		return
// 	}
// 	k = k % n

// 	// 反转整个数组
// 	reverse(nums, 0, n-1)
// 	// 反转前k个元素
// 	reverse(nums, 0, k-1)
// 	// 反转后n-k个元素
// 	reverse(nums, k, n-1)
// }

// func reverse(nums []int, start int, end int) {
// 	for start < end {
// 		nums[start], nums[end] = nums[end], nums[start]
// 		start++
// 		end--
// 	}
// }

// ---使用环状替换的方法---
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n

	count := 0 // 已经移动的元素个数
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]

		// 在一个环内进行替换
		for {
			next := (current + k) % n
			nums[next], prev = prev, nums[next]
			current = next
			count++

			if start == current {
				break
			}
		}
	}
}

func main() {
	// 示例用法
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	// 现在 nums 应该是 [5, 6, 7, 1, 2, 3, 4]
}
