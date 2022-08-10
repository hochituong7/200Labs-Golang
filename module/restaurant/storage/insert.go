package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

// Implement InsertRestaurant in class create_restaurent biz
func (store *sqlStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	if err := store.db.Create(data).Error; err != nil {
		return err

	}
	return nil
}
