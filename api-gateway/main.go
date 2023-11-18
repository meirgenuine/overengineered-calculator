package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", homePageHandler)
	router.GET("/add", operationHandler("add"))
	router.GET("/subtract", operationHandler("subtract"))
	router.GET("/multiply", operationHandler("multiply"))
	router.GET("/divide", operationHandler("divide"))

	router.Run(":8080")
}

func homePageHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Overengineered Calculator! Please use the endpoints /add, /subtract, /multiply, or /divide to perform calculations.",
	})
}

func operationHandler(operation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		num1Str, num2Str := c.Query("num1"), c.Query("num2")
		num1, err1 := strconv.ParseFloat(num1Str, 64)
		num2, err2 := strconv.ParseFloat(num2Str, 64)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid input",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"operation": operation,
			"num1":      num1,
			"num2":      num2,
			"result":    "operation not implemented yet",
		})
	}
}
