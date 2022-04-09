# https://leetcode.com/problems/exchange-seats/submissions/

SELECT (
           CASE
               WHEN MOD(id, 2) = 1 AND id = (SELECT COUNT(*) FROM Seat) THEN id
               WHEN MOD(id, 2) = 1 THEN id + 1
               ELSE id - 1
               END
           ) AS id,
       student
FROM Seat
ORDER BY id;
