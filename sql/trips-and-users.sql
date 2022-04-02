# https://leetcode.com/problems/trips-and-users/submissions/

SELECT request_at `Day`, ROUND(AVG(status != 'completed'), 2) `Cancellation Rate`
FROM Trips t
         JOIN Users u1 ON t.client_id = u1.users_id AND u1.banned = 'No'
         JOIN Users u2 ON t.driver_id = u2.users_id AND u2.banned = 'No'
WHERE request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY request_at;
