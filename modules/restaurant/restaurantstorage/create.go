package restaurantstorage

import (
	"FoodDelivery/common"
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *mysqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil

}
