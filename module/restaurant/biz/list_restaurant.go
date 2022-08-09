package restaurantbiz

import (
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"

	"golang.org/x/net/context"
)

type ListRestaurantStore interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

func NewListRetaurant(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, err
	}
	return result, nil
}
