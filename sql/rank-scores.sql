# https://leetcode.com/problems/rank-scores/submissions/


# sol1
SELECT s1.score, COUNT(DISTINCT s2.score) AS `rank`
FROM Scores AS s1, Scores AS s2
WHERE s1.score <= s2.score
GROUP BY s1.id
ORDER BY `rank`;


# sol2
SELECT score,
       dense_rank() OVER (ORDER BY score DESC) AS 'Rank'
FROM Scores;
