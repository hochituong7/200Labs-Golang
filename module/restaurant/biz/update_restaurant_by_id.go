package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		data *restaurantmodel.RestaurantUpdate,
	) error
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

// method
func (biz *updateRestaurantBiz) UpdateRestaurantById(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {

	//find before update
	//oldData, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"id": id})
	_, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// if oldData.OwnerId == sefl {
	// 	//phân quyền cho update
	// }

	if err := biz.store.UpdateRestaurant(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}
	return nil
}
