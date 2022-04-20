# https://leetcode.com/problems/top-travellers/submissions/

SELECT name, IFNULL(SUM(distance), 0) travelled_distance
FROM Users U
         LEFT JOIN Rides R ON R.user_id = U.id
GROUP BY user_id
ORDER BY travelled_distance DESC, name;
