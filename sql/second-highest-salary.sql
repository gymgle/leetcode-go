# https://leetcode.com/problems/second-highest-salary

# sol1
SELECT (SELECT DISTINCT salary
        FROM Employee
        ORDER BY salary DESC
        LIMIT 1, 1) AS `SecondHighestSalary`;


# sol2
SELECT IFNULL((SELECT DISTINCT salary
               FROM Employee
               ORDER BY salary DESC
               LIMIT 1, 1), NULL) AS `SecondHighestSalary`;
