package restaurantmodel

import (
	"errors"
	"strings"
)

// --- có 3 struct
// 1. main struct (business struct)
// - struct này không dùng cho update đc, vì string ko update chuỗi rỗng đc, nên dùng struct có con trỏ
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

// validate
func (res *RestaurantCreate) Validate() error {
	res.Id = 0 //set giá trị id client truyền lên
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}
