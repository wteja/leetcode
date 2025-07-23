package main

func maximumUniqueSubarray(nums []int) int {
	seen := make(map[int]bool)

	left := 0

	currentSum := 0
	maxSum := 0

	for right := 0; right < len(nums); right++ {
		for seen[nums[right]] {
			delete(seen, nums[left])
			currentSum -= nums[left]
			left++
		}

		seen[nums[right]] = true
		currentSum += nums[right]

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))             // Output: 17
	println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})) // Output: 8
}
