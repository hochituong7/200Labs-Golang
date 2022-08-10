package restaurantgin

import (
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantmodel "food-delivery-service/module/restaurant/model"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateRestaurantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data restaurantmodel.RestaurantUpdate
		id, err := strconv.Atoi(c.Param("restaurant-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//parse data
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(storage)

		if err := biz.UpdateRestaurantById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})

	}

}