package problem1

func twoSum(nums []int, target int) []int {
	pair := []int{}
	set := map[int]int{}
	for i, n := range nums {
		remain := target - n
		if _, ok := set[n]; !ok {
			set[remain] = i
		} else {
			pair = append(pair, set[n])
			pair = append(pair, i)
			break
		}
	}
	return pair
}
