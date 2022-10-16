package controllers

import (
	"net/http"
	"simple-rest-api/initializers"
	"simple-rest-api/models"

	"github.com/gin-gonic/gin"
)

func RestaurantCreate(c *gin.Context) {
	// get data of request
	var input models.Restaurant
	c.ShouldBind(&input)

	// create a restaurant
	restaurant := models.Restaurant{Name: input.Name, Addr: input.Addr}

	result := initializers.DB.Create(&restaurant)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
	}
	// return it
	c.JSON(http.StatusOK, gin.H{
		"restaurant": restaurant,
	})
}

func RestaurantGet(c *gin.Context) {
	// get data
	var data []models.Restaurant
	initializers.DB.Find(&data)
	// do something
	c.JSON(http.StatusOK, gin.H{
		"restaurant": data,
	})
}

func RestaurantGetByID(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	// do something
	var data models.Restaurant
	initializers.DB.Find(&data, id)
	c.JSON(http.StatusOK, gin.H{
		"restaurant": data,
	})
}
