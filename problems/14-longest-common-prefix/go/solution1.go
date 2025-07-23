func longestCommonPrefix(strs []string) string {
	ret := ""

	for i := range len(strs[0]) {
		for _, s := range strs {
			if i >= len(s) || s[i] != strs[0][i] {
				return ret
			}
		}
		ret += string(strs[0][i])
	}

	return ret
}