# coding: utf-8

class Solution(object):
    def findOrder(self, numCourses, prerequisites):
        """
        :type numCourses: int
        :type prerequisites: List[List[int]]
        :rtype: List[int]
        """
        if len(prerequisites) == 0:  # 无依赖关系
            return [i for i in range(numCourses)]

        in_degree = [0] * numCourses  # 入度数组
        adj = [[] for _ in range(numCourses)]  # 依赖邻接表
        for i in prerequisites:
            in_degree[i[0]] += 1
            adj[i[1]].append(i[0])

        queue = list()
        rst = list()
        for i, v in enumerate(in_degree):
            if v == 0:
                queue.append(i)

        while queue:
            c = queue.pop(0)
            rst.append(c)
            for i in adj[c]:
                in_degree[i] -= 1
                if in_degree[i] == 0:
                    queue.append(i)

        if len(rst) != numCourses:
            return []
        return rst


if __name__ == '__main__':
    s = Solution()
    print s.findOrder(4, [[0, 1], [2, 1], [3, 2], [3, 0]])

