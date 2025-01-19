func romanToInt(s string) int {
	roman := map[byte]int{'M': 1000, 'D': 500, 'C': 100, 'L': 50, 'X': 10, 'V': 5, 'I': 1}

	sum := 0

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && roman[s[i]] < roman[s[i+1]] {
			sum -= roman[s[i]]
		} else {
			sum += roman[s[i]]
		}
	}

	return sum
}