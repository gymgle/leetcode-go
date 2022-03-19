class Solution:
    def plusOne(self, digits: list[int]) -> list[int]:
        i = len(digits) - 1
        while i >= 0:
            if digits[i] == 9:
                digits[i] = 0
                i -= 1
            else:
                digits[i] += 1
                break

        if digits[0] == 0:
            digits.insert(0, 1)
        return digits


if __name__ == '__main__':
    s = Solution()
    print(s.plusOne([9, 9]))
