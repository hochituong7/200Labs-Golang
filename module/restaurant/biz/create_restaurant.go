package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

// connect to DB by interface
type CreateRestaurantStore interface {
	InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
	//all func relative IO call orther class then use context
	// return err if insert fail
}

// return pointer struct
func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

// brand struct
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.store.InsertRestaurant(ctx, data); err != nil {
		return err
	}
	return nil
}
