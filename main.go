package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err = runService(db); err != nil {
		log.Fatalln(err)
	}
}

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong ding dong ching chong",
		})
	})

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", func(c *gin.Context) {

			var data Restaurant

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Create(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})

		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id")) //string convert, int

			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			var data Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})

		restaurants.GET("", func(c *gin.Context) {
			var data []Restaurant //get list nen phai dùng mảng

			type Filter struct { //filter them city_id
				CityId int `json:"city_id" form:"city_id"` //de nhan gia tri luon thi speard string thi dung form
			}

			var filter Filter

			c.ShouldBind(&filter) // bind data filter vao CityId

			newDB := db

			if filter.CityId > 0 { //TH co dung filter (default cua CityId = 0, != ng hia la co ton tai)
				newDB = db.Where("city_id = ?", filter.CityId) //Neu co ton tai filter thi where, ko thi giu nguyen newDB = db
			}

			if err := newDB.Find(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, data)
		})

		restaurants.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			var data Restaurant

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
			}
		})

		restaurants.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Table(Restaurant{}.TableName()).
				Where("id = ?", id).
				Delete(nil).Error; err != nil { //delete whole thing
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
