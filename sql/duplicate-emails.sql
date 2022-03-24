# https://leetcode.com/problems/duplicate-emails/submissions/

SELECT email
FROM Person
GROUP BY email
HAVING COUNT(*) > 1;
