func longestCommonPrefix(strs []string) string {
    ret := ""

	for i in range(len(strs[0])) {
		for j in range(1, len(strs)) {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return ret
			}
		}
		ret += string(strs[0][i])
	}

	return ret
}