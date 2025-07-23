package main

func makeFancyString(s string) string {
	var result []rune
	var count = 1

	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			result = append(result, rune(s[i]))
		}
	}

	return string(result)
}

func main() {
	println(makeFancyString("leeetcode")) // Output: "leetcode"
	println(makeFancyString("aaabaaaa"))  // Output: "aabaa"
	println(makeFancyString("aab"))       // Output: "aab"
}
