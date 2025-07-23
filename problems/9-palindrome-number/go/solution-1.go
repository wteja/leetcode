package problem9

import "strconv"

func isPalindrome(x int) bool {
	isPalindrom := true
	str := strconv.Itoa(x)
	arr := []rune(str)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			isPalindrom = false
			break
		}
	}
	return isPalindrom
}
