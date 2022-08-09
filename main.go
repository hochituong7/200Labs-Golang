package main

import (
	"log"

	restaurantgin "food-delivery-service/module/restaurant/transport/gin"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	///dsn := os.Getenv("MYSQL_CONNECTION")
	dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MYSQL: ", err)
	}

	log.Println("Connected: ", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(db))
			restaurants.GET("/:restaurant-id", restaurantgin.GetRestaurantHandler(db))
			restaurants.GET("", restaurantgin.ListRestaurant(db))
			restaurants.PUT("/:restaurant-id", restaurantgin.UpdateRestaurantHandler(db))
			restaurants.DELETE("/:restaurant-id", restaurantgin.DeleteRestaurantHandler(db))
		}
	}

	router.Run(":3003") //default 8080

}
