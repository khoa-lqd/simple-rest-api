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
	// get id from url
	id := c.Param("id")
	// do something
	var data models.Restaurant
	initializers.DB.Find(&data, id)

	c.JSON(http.StatusOK, gin.H{
		"restaurant": data,
	})
}

func RestaurantUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data off request
	var updated struct {
		Name string
		Addr string
	}
	c.ShouldBind(&updated)
	// find the data were updating
	var data models.Restaurant
	initializers.DB.First(&data, id)
	// update the data
	initializers.DB.Model(&data).Updates(models.Restaurant{Name: updated.Name, Addr: updated.Addr})
	// respond with it
	c.JSON(http.StatusOK, gin.H{
		"restaurant": updated,
	})
}

func RestaurantDelete(c *gin.Context) {
	id := c.Param("id")
	var delete models.Restaurant

	initializers.DB.Delete(&models.Restaurant{}, id)

	c.JSON(http.StatusOK, gin.H{
		"restaurant": delete,
	})

}
