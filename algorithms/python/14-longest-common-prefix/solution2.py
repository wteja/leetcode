class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        
        for i, letters in enumerate(zip(*strs)):
            if len(set(letters)) > 1:
                return strs[0][:i]
        
        else:
            return min(strs)

# Zip each characters and check the difference with set, I found difference just stop and return.