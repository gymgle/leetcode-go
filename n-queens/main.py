class Solution:
    def solveNQueens(self, n: int) -> list[list[str]]:
        def trace(r: int, bd: list[list[str]]):
            if r == n:
                answer = list()
                for r in range(n):
                    answer.append(''.join(bd[r]))
                solutions.append(answer)
                return
            for c in range(n):
                if c in columns or (r - c) in diagonal1 or (r + c) in diagonal2:
                    continue
                bd[r][c] = 'Q'
                columns.add(c)
                diagonal1.add(r - c)
                diagonal2.add(r + c)
                trace(r + 1, bd)
                bd[r][c] = '.'
                columns.remove(c)
                diagonal1.remove(r - c)
                diagonal2.remove(r + c)

        solutions = list()
        board = list()
        for i in range(n):
            row = ['.'] * n
            board.append(row)

        columns = set()
        diagonal1 = set()
        diagonal2 = set()
        trace(0, board)
        return solutions


if __name__ == '__main__':
    s = Solution()
    print(s.solveNQueens(4))
