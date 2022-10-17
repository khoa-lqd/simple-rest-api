package main

import (
	"simple-rest-api/initializers"
	"simple-rest-api/models"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Restaurant{})
}
