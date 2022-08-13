package restaurantstorage

import (
	"context"
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

func (store *sqlStore) DeleteRestaurant(
	ctx context.Context,
	cond map[string]interface{},
) error {

	//data truyền vào đã là con trỏ nên truyển vào Updates ko cần nữa
	if err := store.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where(cond).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
