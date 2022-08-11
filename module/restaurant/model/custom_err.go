package restaurantmodel

import "food-delivery-service/common"

// custom error this model
var (
	ErrNameCannotBeBlank = common.NewCustomError(nil, "restaurant name can't be blank", "ErrNameCannotBeBlank")
)
