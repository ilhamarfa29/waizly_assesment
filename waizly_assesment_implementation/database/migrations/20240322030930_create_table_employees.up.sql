CREATE TABLE employees (
	employee_id varchar(100) PRIMARY KEY,
	employee_name varchar(255) NOT NULL,
	job_title varchar(50),
	salary decimal,
	department varchar(50),
	joined_date timestamp(0) with time zone DEFAULT CURRENT_TIMESTAMP(6)
);