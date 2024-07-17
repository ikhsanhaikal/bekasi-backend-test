-- name: SoalDua :many
select * from employees e
where e.terminationdate isnull
and lastname = 'smith'
order by
	lastname asc,
	firstname asc;

-- name: SoalTiga :many
select e.*
from 
  employees e
  left join annualreviews a on a.empid = e.id
  where empid isnull 
order by 
  e.hiredate ;

-- name: SoalEmpat :one
select max(e.hiredate) as terbaru, min(e.hiredate) as tua,
 MAX(e.hiredate)- MIN(e.hiredate) AS DateDifference  
from employees e; 

-- name: SoalLimaCountTotalUlasan :many
select e.id, count(a.empid) as total 
from 
  employees e
  left join annualreviews a on a.empid = e.id 
GROUP BY 
  e.id ;

-- name: SoalLima :many
select * from employees e;


