package main

func maximumGain(s string, x int, y int) int {
	first, second := "ab", "ba"
	firstScore, secondScore := x, y

	if y > x {
		first, second = "ba", "ab"
		firstScore, secondScore = y, x
	}

	removeStr := func(str string, pattern string) (string, int) {
		stack := []byte{}
		count := 0

		for i := range len(str) {
			stack = append(stack, str[i])

			if len(stack) >= 2 &&
				stack[len(stack)-2] == pattern[0] &&
				stack[len(stack)-1] == pattern[1] {
				stack = stack[:len(stack)-2]
				count++
			}
		}

		return string(stack), count
	}

	remaining, count1 := removeStr(s, first)
	_, count2 := removeStr(remaining, second)

	return (count1 * firstScore) + (count2 * secondScore)
}

func main() {
	println(maximumGain("cdbcbbaaabab", 4, 5))   // Output: 19
	println(maximumGain("aabbaaxybbaabb", 5, 4)) // Output: 20
}
