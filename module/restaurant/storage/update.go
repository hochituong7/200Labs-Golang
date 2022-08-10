package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

func (store *sqlStore) UpdateRestaurant(
	ctx context.Context,
	cond map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {

	//data truyền vào đã là con trỏ nên truyển vào Updates ko cần nữa
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
