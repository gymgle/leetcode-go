class Solution:
    def letterCombinations(self, digits: str) -> list[str]:
        def trace(index: int, combine: str):
            if index == n:
                solutions.append(combine)
                return
            for i in keyboard[digits[index]]:
                combine += i
                trace(index + 1, combine)
                combine = combine[:-1]

        if not digits:
            return []
        keyboard = {
            '2': ['a', 'b', 'c'],
            '3': ['d', 'e', 'f'],
            '4': ['g', 'h', 'i'],
            '5': ['j', 'k', 'l'],
            '6': ['m', 'n', 'o'],
            '7': ['p', 'q', 'r', 's'],
            '8': ['t', 'u', 'v'],
            '9': ['w', 'x', 'y', 'z']
        }
        solutions = list()
        combine = ''
        n = len(digits)
        trace(0, combine)
        return solutions


if __name__ == '__main__':
    s = Solution()
    print(s.letterCombinations('23'))
