SELECT class 
FROM Courses
GROUP BY class
Having count(distinct student)>=5

