CREATE TABLE IF NOT EXISTS annualreviews (
  id   serial PRIMARY KEY,
	empId serial NOT NULL,
	reviewdate date NOT NULL DEFAULT CURRENT_DATE
);

-- CREATE TABLE public.annualreviews (
-- 	id serial NOT NULL,
-- 	empid serial NOT NULL,
-- 	reviewdate date NOT NULL,
-- 	CONSTRAINT annualreviews_pk PRIMARY KEY (id),
-- 	CONSTRAINT annualreviews_employees_fk FOREIGN KEY (empid) REFERENCES public.employees(id)
-- );
