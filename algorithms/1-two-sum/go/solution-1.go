package main

import "fmt"

func twoSum(nums []int, target int) []int {
	pair := []int{}
	set := map[int]int{}
	for i, n := range nums {
		if _, ok := set[n]; !ok {
			set[target-n] = i
		} else {
			pair = append(pair, set[n])
			pair = append(pair, i)
			break
		}
	}
	return pair
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // Output: [1, 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // Output: [0, 1]
}
