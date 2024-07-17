CREATE TABLE IF NOT EXISTS employees (
  id   serial PRIMARY KEY,
  firstname varchar (100)      NOT NULL,
  lastname varchar (100)  NOT NULL,
	hiredate date NOT NULL,
	terminationdate date,
	salary numeric 
);

-- CREATE TABLE public.employees (
-- 	id serial NOT NULL,
-- 	firstname varchar NULL,
-- 	lastname varchar NULL,
-- 	hiredate date NOT NULL,
-- 	terminationdate date NULL,
-- 	salary numeric NULL,
-- 	CONSTRAINT employees_pk PRIMARY KEY (id)
-- );

-- emmployee
-- bob smith 2009-06-20 2016-01-01 10000
-- joe jarrot 2010-02-12 null 20000
-- nancy soley 2012-03-14 null 30000
-- keith widjaja 2013-09-10 2014-01-01 20000
-- kelly smalls 2013-09-10 null 20000
-- frank nguyen 2015-04-10 2015-05-01 60000

