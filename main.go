package main

import (
	"fmt"
	"simple-rest-api/controllers"
	"simple-rest-api/initializers"

	bugsnaggin "github.com/bugsnag/bugsnag-go-gin"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gin-gonic/gin"
)

func main() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	router := gin.Default()

	router.Use(bugsnaggin.AutoNotify(bugsnag.Configuration{
		// Your Bugsnag project API key, required unless set as environment
		// variable $BUGSNAG_API_KEY
		APIKey: "82789214141a730634bce15cb236141f",
		// The import paths for the Go packages containing your source files
	}))

	bugsnag.Notify(fmt.Errorf("Test error"))

	router.POST("/restaurants", controllers.RestaurantCreate)

	router.GET("/restaurants", controllers.RestaurantGet)

	router.GET("/restaurants/:id", controllers.RestaurantGetByID)

	router.PUT("/restaurants/:id", controllers.RestaurantUpdate)

	router.DELETE("/restaurants/:id", controllers.RestaurantDelete)

	if err := router.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		fmt.Println(err)
	}
}
