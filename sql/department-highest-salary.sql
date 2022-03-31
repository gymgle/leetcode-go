# https://leetcode.com/problems/department-highest-salary/

SELECT D.name AS `Department`, E.name AS `Employee`, E.salary AS `Salary`
FROM Employee E,
     Department D
WHERE E.departmentId = D.id
  AND (E.departmentId, E.salary) IN (SELECT departmentId, MAX(salary) FROM Employee GROUP BY departmentId);
