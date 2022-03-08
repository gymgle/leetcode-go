class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ''

        def get_common_prefix(s1, s2):
            n = min(len(s1), len(s2))
            common_prefix = ''
            for i in range(n):
                if s1[i] != s2[i]:
                    break
                common_prefix += s1[i]
            return common_prefix

        common = strs[0]
        for i in strs[1:]:
            common = get_common_prefix(common, i)

        return common
