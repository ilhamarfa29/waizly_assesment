package main

import (
	"fmt"

	"example.com/waizly-test-implementation/database"
	emp "example.com/waizly-test-implementation/internal/app/employees/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/employees/:id", emp.ReadEmployee)
	r.GET("/employees", emp.ReadEmployees)
	r.POST("/employees", emp.CreateEmployee)
	r.PUT("/employees/:id", emp.UpdateEmployee)
	r.DELETE("/employees/:id", emp.DeleteEmployee)
	r.Run(":5000")
}
