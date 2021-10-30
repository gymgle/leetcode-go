"""
最佳思路：通过哈希存储映射，从右到左遍历每个字符，检查当前字符的映射值，是否大于上一个，是则加当前的值，否则，减去当前的值。
"""


class Solution:
    def romanToInt(self, s: str) -> int:
        m = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }
        num = 0
        i, n = 0, len(s)
        while i < n:
            if s[i] == 'I' and i + 1 < n and s[i + 1] in ['V', 'X']:
                num += (int(m[s[i + 1]]) - int(m[s[i]]))
                i += 2
            elif s[i] == 'X' and i + 1 < n and s[i + 1] in ['L', 'C']:
                num += (int(m[s[i + 1]]) - int(m[s[i]]))
                i += 2
            elif s[i] == 'C' and i + 1 < n and s[i + 1] in ['D', 'M']:
                num += (int(m[s[i + 1]]) - int(m[s[i]]))
                i += 2
            else:
                num += int(m[s[i]])
                i += 1
        return num


if __name__ == '__main__':
    sol = Solution()
    print(sol.romanToInt('LVIII'))  # 58
