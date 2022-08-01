package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("MYSQL_CONNECTION")
	//dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MYSQL: ", err)
	}
	log.Println("Connected: ", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		restaurant := v1.Group("/restaurants")
		restaurant.POST("/restaurants", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})
		// v1.POST("/items", createItem(db))           // create item
		// v1.GET("/items", getListOfItems(db))        // list items
		// v1.GET("/items/:id", readItemById(db))      // get an item by ID
		// v1.PUT("/items/:id", editItemById(db))      // edit an item by ID
		// v1.DELETE("/items/:id", deleteItemById(db)) // delete an item by ID
	}

	router.Run(":3003") //default 8080

}

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}
