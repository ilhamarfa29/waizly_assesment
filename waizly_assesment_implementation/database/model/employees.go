package database

import "time"

type Employees struct {
	EmployeeId   string    `json:"employee_id"`
	EmployeeName string    `json:"employee_name"`
	JobTitle     string    `json:"job_title"`
	Salary       float64   `json:"salary"`
	Department   string    `json:"department"`
	JoinedDate   time.Time `json:"joined_date"`
}
