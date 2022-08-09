package restaurantgin

import (
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getListRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_ = paging.Validate()
		store := restaurantmodel

		c.JSON(http.StatusOK, gin.H{"paging": paging, "data": result})
	}

}
