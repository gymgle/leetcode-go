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
