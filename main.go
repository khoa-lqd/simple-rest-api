package main

import (
	"fmt"
	"simple-rest-api/controllers"
	"simple-rest-api/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	router := gin.Default()
	router.POST("/restaurants", controllers.RestaurantCreate)

	router.GET("/restaurants", controllers.RestaurantGet)

	router.GET("/restaurants/:id", controllers.RestaurantGetByID)

	router.PUT("/restaurants/:id", controllers.RestaurantUpdate)

	router.DELETE("/restaurants/:id", controllers.RestaurantDelete)

	if err := router.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		fmt.Println(err)
	}
}
