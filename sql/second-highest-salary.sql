# https://leetcode.com/problems/second-highest-salary

SELECT (SELECT DISTINCT salary
        FROM Employee
        ORDER BY salary DESC
        LIMIT 1, 1) AS `SecondHighestSalary`;
