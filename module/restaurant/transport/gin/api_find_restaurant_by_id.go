package restaurantgin

import (
	"food-delivery-service/common"
	"food-delivery-service/components"
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurantHandler(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("restaurant-id"))

		uid, err := common.FromBase58(c.Param("restaurant-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		storage := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewFindRestaurantBiz(storage)

		data, err := biz.FindRestaurantById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.Mask(true)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}

}
