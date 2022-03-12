class Solution:
    def isValid(self, s: str) -> bool:
        m = {
            '(': ')',
            '{': '}',
            '[': ']'
        }
        if not s:
            return False
        if len(s) % 2 != 0:
            return False

        stack = []
        for i in s:
            if m.get(i, ''):
                stack.append(i)
            elif not stack or m[stack[-1]] != i:
                return False
            else:
                stack.pop()
        return True if not stack else False
