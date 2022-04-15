# https://leetcode.com/problems/market-analysis-i/

SELECT user_id AS `buyer_id`, join_date, COUNT(order_id) AS `orders_in_2019`
FROM Users s
         LEFT JOIN Orders o ON s.user_id = o.buyer_id AND order_date BETWEEN '2019-01-01' AND '2019-12-31'
GROUP BY user_id;
