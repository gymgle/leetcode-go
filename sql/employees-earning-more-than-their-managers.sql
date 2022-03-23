# https://leetcode.com/problems/employees-earning-more-than-their-managers/

# sol 1
SELECT a.name AS `Employee`
FROM Employee AS a,
     Employee AS b
WHERE a.managerId = b.id
  AND a.salary > b.salary;


# sol 2
SELECT a.name AS `Employee`
FROM Employee AS a
         JOIN Employee AS b ON a.managerId = b.id AND a.salary > b.salary;