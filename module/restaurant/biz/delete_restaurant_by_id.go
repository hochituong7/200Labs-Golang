package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	DeleteRestaurant(
		ctx context.Context,
		cond map[string]interface{},
	) error
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

// method
func (biz *deleteRestaurantBiz) DeleteRestaurantById(ctx context.Context, id int) error {

	//find before delete
	//oldData, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"id": id})
	_, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	// if oldData.OwnerId == sefl {
	// 	//phân quyền cho delete
	// }

	if err := biz.store.DeleteRestaurant(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
