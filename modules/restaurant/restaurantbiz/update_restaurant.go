package restaurantbiz

import (
	"FoodDelivery/common"
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"context"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)

	UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}
	return nil

}
