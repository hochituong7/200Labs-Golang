package main

import (
	"log"

	"food-delivery-service/components"
	"food-delivery-service/middleware"

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
	router.Use(middleware.Recover()) //show ra các lỗi khi bị crash, nếu ko có thì ko show gì (500 internal error with postman)
	appCtx := components.NewAppContext(db)

	mainRoute(router, appCtx)
	router.Run(":3003") //default 8080

}
