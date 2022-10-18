package restaurantgin

import (
	"net/http"
	restaurantbiz "simple-rest-api/modules/restaurant/biz"
	restaurantmodel "simple-rest-api/modules/restaurant/model"
	restaurantstorage "simple-rest-api/modules/restaurant/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input restaurantmodel.RestaurantCreate

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// return it
		c.JSON(http.StatusOK, gin.H{
			"restaurant": input,
		})
	}
}
