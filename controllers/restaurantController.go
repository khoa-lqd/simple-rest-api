package controllers

import (
	"net/http"
	"simple-rest-api/initializers"
	"simple-rest-api/models"

	"github.com/gin-gonic/gin"
)

func RestaurantCreate(c *gin.Context) {
	// get data of request

	// validate input
	var input models.Restaurant

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	// do something
	var data models.Restaurant

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "restaurant not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"restaurant": data,
	})
}

func RestaurantUpdate(c *gin.Context) {
	// find the data were updating

	var data models.Restaurant

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "restaurant not found",
		})
		return
	}

	// get data off request

	var updated struct {
		Name string
		Addr string
	}
	// validate input

	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update the data
	initializers.DB.Model(&data).Updates(models.Restaurant{Name: updated.Name, Addr: updated.Addr})
	// respond with it
	c.JSON(http.StatusOK, gin.H{
		"restaurant": updated,
	})
}

func RestaurantDelete(c *gin.Context) {
	// id := c.Param("id")
	var delete models.Restaurant
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&delete).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	initializers.DB.Delete(&delete)

	c.JSON(http.StatusOK, gin.H{
		"restaurant": true,
	})

}
