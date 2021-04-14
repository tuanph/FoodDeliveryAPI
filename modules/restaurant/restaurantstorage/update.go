package restaurantstorage

import (
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"context"
)

func (s *mysqlStore) UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
