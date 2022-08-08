package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
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
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", createRestaurant(db))
			restaurants.GET("/:restaurant-id", getRestaurant(db))
			restaurants.GET("", getListRestaurant(db))
			restaurants.PUT("/:restaurant-id", updateRestaurant(db))
			restaurants.DELETE("/:restaurant-id", deleteRestaurant(db))
		}
	}

	router.Run(":3003") //default 8080

}

//--- có 3 struct
//1. main struct (business struct)
//- struct này không dùng cho update đc, vì string ko update chuỗi rỗng đc, nên dùng struct có con trỏ
type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"` //tag
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

// khi truy vấn data thì dùng bảng khai báo này để biết bảng nào
func (Restaurant) TableName() string {
	return "restaurants"
}

// 2.truct thao tác data với db, dùng con trỏ cho update chuỗi rỗng
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

// 3.truct thao tác data với db
// struct chỉ lấy những field cần thiết
type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;"` // id = "-" không nhận params id từ client lên
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

//validate
func (res *RestaurantCreate) Validate() error {
	res.Id = 0 //set giá trị id client truyền lên
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}

//vieestn hàm create riêng
func createRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"ok": 1})
		var data RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//validate
		if err := data.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// insert db
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}

}

func getRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"ok": 1})
		var data Restaurant

		id, err := strconv.Atoi(c.Param("restaurant-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}

}

//hello architecture
func getListRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"ok": 1})
		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 5
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []Restaurant
		if err := db.Table(Restaurant{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Limit(paging.Limit).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"paging": paging, "data": result})
	}

}

func updateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"ok": 1})
		var data RestaurantUpdate

		id, err := strconv.Atoi(c.Param("restaurant-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}

}

func deleteRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"ok": 1})

		id, err := strconv.Atoi(c.Param("restaurant-id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}

}
