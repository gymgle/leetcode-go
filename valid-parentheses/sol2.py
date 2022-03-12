from operator import contains


class Solution:
    def isValid(self, s: str) -> bool:
        if not s:
            return False
        if len(s) % 2 != 0:
            return False
        while True:
            if contains(s, '[]'):
                s = s.replace('[]', '')
            elif contains(s, '{}'):
                s = s.replace('{}', '')
            elif contains(s, '()'):
                s = s.replace('()', '')
            else:
                break
        return True if not s else False
