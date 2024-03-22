package main

import (
	"fmt"

	"example.com/waizly-test-implementation/controllers"
	"example.com/waizly-test-implementation/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/employees/:id", controllers.ReadEmployee)
	r.GET("/employees", controllers.ReadEmployees)
	r.POST("/employees", controllers.CreateEmployee)
	r.PUT("/employees/:id", controllers.UpdateEmployee)
	r.DELETE("/employees/:id", controllers.DeleteEmployee)
	r.Run(":5000")
}
