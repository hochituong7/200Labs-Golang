package main

import (
	"food-delivery-service/components"
	restaurantgin "food-delivery-service/module/restaurant/transport/gin"

	"github.com/gin-gonic/gin"
)

func mainRoute(router *gin.Engine, appCtx components.AppContext) {
	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(appCtx))
			restaurants.GET("/:restaurant-id", restaurantgin.GetRestaurantHandler(appCtx))
			restaurants.GET("", restaurantgin.ListRestaurant(appCtx))
			restaurants.PUT("/:restaurant-id", restaurantgin.UpdateRestaurantHandler(appCtx))
			restaurants.DELETE("/:restaurant-id", restaurantgin.DeleteRestaurantHandler(appCtx))
		}
	}
}
