package restaurantgin

import (
	"food-delivery-service/common"
	"food-delivery-service/components"
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantmodel "food-delivery-service/module/restaurant/model"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		// go func() {
		// 	defer common.Recover() //nếu mở go func thì nên để thêm hàm recover tránh bị crash

		// 	var arr []int
		// 	fmt.Println(arr[0])
		// }()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			// c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			// return
			panic(err)
		}

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		_ = paging.Validate()
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRetaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i := range result {
			result[i].Mask(true)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}

}
