# https://leetcode.com/problems/tree-node/

SELECT id,
       CASE
           WHEN t.p_id IS NULL THEN 'Root'
           WHEN t.id IN (SELECT p_id FROM Tree) THEN 'Inner'
           ELSE 'Leaf'
           END AS `Type`
FROM Tree t;
