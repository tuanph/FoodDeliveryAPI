package restaurantbiz

import (
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"context"
	"errors"
)

//Declare Interface
type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

//Implemet methods
func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("restaurant name can't be blank")
	}
	err := biz.store.Create(ctx, data)
	return err
}
