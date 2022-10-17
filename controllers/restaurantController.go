package controllers

import (
	"net/http"
	"simple-rest-api/initializers"
	"simple-rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RestaurantCreate(c *gin.Context) {
	// 1. get data of request
	// validate input
	var input models.RestaurantCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 2.create a restaurant
	restaurant := models.RestaurantCreate{Name: input.Name, Addr: input.Addr}

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
	id, _ := strconv.Atoi(c.Param("id"))

	// do something
	var data models.Restaurant
	if err := initializers.DB.Where("id = ?", id).First(&data).Error; err != nil {
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

	var (
		data   models.Restaurant
		update models.RestaurantUpdate
	)

	// get data off request
	// validate input
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validate input
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "restaurant not found",
		})
		return
	} // no need but its ok to have

	// update the data
	initializers.DB.Model(&data).Updates(models.Restaurant{Name: update.Name, Addr: update.Addr})

	// respond with it
	c.JSON(http.StatusOK, gin.H{
		"restaurant": data,
	})
}

func RestaurantDelete(c *gin.Context) {
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
