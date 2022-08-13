package restaurantmodel

import (
	"food-delivery-service/common"
	"strings"
)

const EntityName = "Restaurant"

// --- có 3 struct
// 1. main struct (business struct)
// - struct này không dùng cho update đc, vì string ko update chuỗi rỗng đc, nên dùng struct có con trỏ
type Restaurant struct {
	common.SQLModel
	OwnerId int    `json:"owner_id" gorm:"column:owner_id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Addr    string `json:"address" gorm:"column:addr;"`
}

func (r *Restaurant) Mask(isAdminOwner bool) {
	r.SQLModel.Mask(common.DbTypeRestaurant)
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
	common.SQLModel
	//Id   int    `json:"id" gorm:"column:id;"` // id = "-" không nhận params id từ client lên
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
		return ErrNameCannotBeBlank
	}

	return nil
}
