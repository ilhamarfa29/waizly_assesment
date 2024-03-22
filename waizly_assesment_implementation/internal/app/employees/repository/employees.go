package employee

import (
	"errors"
	"time"

	"example.com/waizly-test-implementation/database"
	modelDb "example.com/waizly-test-implementation/database/model"
	"github.com/google/uuid"
)

func CreateEmployeeRepo(employee *modelDb.Employees) (*modelDb.Employees, error) {

	id := uuid.New()
	employee.EmployeeId = id.String()
	if employee.JoinedDate == nil {
		t := time.Now()
		employee.JoinedDate = &t
	}

	res := database.DB.Create(employee)
	if res.RowsAffected == 0 {
		return nil, errors.New("error creating a employee")
	}

	return employee, nil
}

func ReadEmployeeRepo(id string) (*modelDb.Employees, error) {

	var employee modelDb.Employees

	res := database.DB.Where("employee_id = ?", id).Find(&employee)
	if res.RowsAffected == 0 {
		return nil, errors.New("employee not found")
	}

	return &employee, nil
}

func ReadEmployeesRepo() ([]modelDb.Employees, error) {

	var employees []modelDb.Employees
	res := database.DB.Find(&employees)
	if res.Error != nil {
		return nil, errors.New("employee not found")
	}

	return employees, nil
}

func UpdateEmployeeRepo(employee *modelDb.Employees) (*modelDb.Employees, error) {
	var updateEmployee modelDb.Employees
	res := database.DB.Model(&updateEmployee).Where("employee_id = ?", employee.EmployeeId).Updates(employee)

	if res.RowsAffected == 0 {
		return nil, errors.New("employee not updated")
	}

	return employee, nil
}

func DeleteEmployeeRepo(id string) error {
	var employee modelDb.Employees
	res := database.DB.Where("employee_id = ?", id).Find(&employee)
	if res.RowsAffected == 0 {
		return errors.New("employee not found")
	}
	database.DB.Delete(&employee, "employee_id = ?", id)

	return nil
}
