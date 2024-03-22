package employee

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	modelDb "example.com/waizly-test-implementation/database/model"
	repo "example.com/waizly-test-implementation/internal/app/employees/repository"
	constants "example.com/waizly-test-implementation/internal/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	domainNameEmployee = "employee"
)

func initLog() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	return log
}

func CreateEmployee(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.CreateProcess, domainNameEmployee))
	defer log.Info(fmt.Printf(constants.CreateProcessDone, domainNameEmployee))

	var employee *modelDb.Employees
	err := c.ShouldBind(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		log.Error(err.Error())
		return
	}

	id := uuid.New()
	employee.EmployeeId = id.String()

	result, err := repo.CreateEmployeeRepo(employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a employee",
		})
		log.Error("error creating a employee")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employee": result,
	})

	log.Info(fmt.Printf(constants.CreateProcessSuccess, domainNameEmployee))

	return
}

func ReadEmployee(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.ReadByIdProcess, domainNameEmployee))
	defer log.Info(fmt.Printf(constants.ReadByIdProcessDone, domainNameEmployee))

	id := c.Param("id")
	result, err := repo.ReadEmployeeRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})

		log.Error("employee not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employee": result,
	})
	log.Info(fmt.Printf(constants.ReadByIdProcessSuccess, domainNameEmployee))

	return
}

func ReadEmployees(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.ReadProcess, domainNameEmployee))
	defer log.Info(fmt.Printf(constants.ReadProcessDone, domainNameEmployee))

	result, err := repo.ReadEmployeesRepo()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("employees not found"),
		})

		log.Error("employees not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employees": result,
	})
	log.Info(fmt.Printf(constants.ReadProcessSuccess, domainNameEmployee))

	return
}

func UpdateEmployee(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.UpdateProcess, domainNameEmployee))
	defer log.Info(fmt.Printf(constants.UpdateProcessDone, domainNameEmployee))

	var employee modelDb.Employees
	id := c.Param("id")
	err := c.ShouldBind(&employee)
	employee.EmployeeId = id

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		log.Error(err.Error())
		return
	}

	result, err := repo.UpdateEmployeeRepo(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		log.Error("employee not updated")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"employee": result,
	})
	log.Info(fmt.Printf(constants.UpdateProcessSuccess, domainNameEmployee))

	return
}

func DeleteEmployee(c *gin.Context) {
	log := initLog()
	log.Info(fmt.Printf(constants.DeleteProcess, domainNameEmployee))
	defer log.Info(fmt.Printf(constants.DeleteProcessDone, domainNameEmployee))

	id := c.Param("id")
	_, err := repo.ReadEmployeeRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})

		log.Error("employee not found")
		return
	}

	err = repo.DeleteEmployeeRepo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		log.Error(fmt.Printf(constants.ErrDataNotFound, domainNameEmployee))

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "employee deleted successfully",
	})
	log.Info(fmt.Printf(constants.DeleteProcessSuccess, domainNameEmployee))

	return
}
