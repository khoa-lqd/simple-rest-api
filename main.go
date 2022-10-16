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

	// restaurants := r.Group("/restaurants")
	// {
	// 	restaurants.POST("", func(c *gin.Context) {

	// 		var data Restaurant

	// 		if err := c.ShouldBind(&data); err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		if err := db.Create(&data).Error; err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}
	// 		c.JSON(http.StatusOK, data)
	// 	})

	router.GET("/restaurants", controllers.RestaurantGet)

	// 	restaurants.GET("/:id", func(c *gin.Context) {
	// 		id, err := strconv.Atoi(c.Param("id")) //string convert, int

	// 		if err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}
	// 		var data Restaurant

	// 		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}
	// 		c.JSON(http.StatusOK, data)
	// 	})
	router.GET("/restaurants/:id", controllers.RestaurantGetByID)

	// 	restaurants.GET("", func(c *gin.Context) {
	// 		var data []Restaurant //get list nen phai dùng mảng

	// 		type Filter struct { //filter them city_id
	// 			CityId int `json:"city_id" form:"city_id"` //de nhan gia tri luon thi speard string thi dung form
	// 		}

	// 		var filter Filter

	// 		c.ShouldBind(&filter) // bind data filter vao CityId

	// 		newDB := db

	// 		if filter.CityId > 0 { //TH co dung filter (default cua CityId = 0, != ng hia la co ton tai)
	// 			newDB = db.Where("city_id = ?", filter.CityId) //Neu co ton tai filter thi where, ko thi giu nguyen newDB = db
	// 		}

	// 		if err := newDB.Find(&data).Error; err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}
	// 		c.JSON(http.StatusOK, data)
	// 	})

	// 	restaurants.PATCH("/:id", func(c *gin.Context) {
	// 		id, err := strconv.Atoi(c.Param("id"))

	// 		if err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		var data Restaurant

	// 		if err := c.ShouldBind(&data); err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 		}
	// 	})

	// 	restaurants.DELETE("/:id", func(c *gin.Context) {
	// 		id, err := strconv.Atoi(c.Param("id"))

	// 		if err != nil {
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}

	// 		if err := db.Table(Restaurant{}.TableName()).
	// 			Where("id = ?", id).
	// 			Delete(nil).Error; err != nil { //delete whole thing
	// 			c.JSON(401, gin.H{
	// 				"error": err.Error(),
	// 			})
	// 			return
	// 		}
	// 		c.JSON(http.StatusOK, gin.H{"ok": 1})
	// 	})
	// }

	if err := router.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		fmt.Println(err)
	}
}

// func (Restaurant) TableName() string {
// 	return "restaurants"
// }
