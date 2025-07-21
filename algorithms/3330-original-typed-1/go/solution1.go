package main

func possibleStringCount(word string) int {
	versions := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			versions++
		}
	}
	return versions
}

func main() {
	println(possibleStringCount("abbcccc")) // Output: 5
	println(possibleStringCount("abcd"))    // Output: 1
	println(possibleStringCount("aaaa"))    // Output: 4
}
