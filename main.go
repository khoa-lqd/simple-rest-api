package main

import (
	"log"
	"net/http"

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
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
