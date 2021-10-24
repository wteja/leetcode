class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        dict = { '(': ')', '[': ']', '{': '}' }
        for char in s:
            if char in dict.keys():
                stack.append(dict[char])
            elif char in dict.values():
                if stack == [] or stack.pop() != char:
                    return False
            else:
                return False
            
        return stack == []