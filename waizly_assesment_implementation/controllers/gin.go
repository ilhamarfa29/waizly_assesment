package controllers

import (
	"errors"
	"net/http"

	"example.com/waizly-test-implementation/database"
	modelDb "example.com/waizly-test-implementation/database/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateEmployee(c *gin.Context) {
	var employee *modelDb.Employees
	err := c.ShouldBind(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	id := uuid.New()
	employee.EmployeeId = id.String()

	res := database.DB.Create(employee)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a employee",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
	return
}

func ReadEmployee(c *gin.Context) {
	var employee modelDb.Employees
	id := c.Param("id")
	res := database.DB.Where("employee_id = ?", id).Find(&employee)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
	return
}

func ReadEmployees(c *gin.Context) {
	var employees []modelDb.Employees
	res := database.DB.Find(&employees)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("employees not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employees": employees,
	})
	return
}

func UpdateEmployee(c *gin.Context) {
	var employee modelDb.Employees
	id := c.Param("id")
	err := c.ShouldBind(&employee)
	employee.EmployeeId = id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateEmployee modelDb.Employees
	res := database.DB.Model(&updateEmployee).Where("employee_id = ?", id).Updates(employee)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "employee not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": employee,
	})
	return
}

func DeleteEmployee(c *gin.Context) {
	var employee modelDb.Employees
	id := c.Param("id")
	res := database.DB.Where("employee_id = ?", id).Find(&employee)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}
	database.DB.Delete(&employee, "employee_id = ?", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "employee deleted successfully",
	})
	return
}
