package restaurantbiz

import (
	"FoodDelivery/common"
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"context"
	"log"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}
type RestaurantLikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}
type listRestaurantBiz struct {
	store               ListRestaurantStore
	restaurantLikestore RestaurantLikeStore
}

func NewListRestaurantBiz(store ListRestaurantStore, restaurantLikestore RestaurantLikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, restaurantLikestore: restaurantLikestore}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	ids := make([]int, len(result))
	for i := range result {
		ids[i] = result[i].Id
	}

	mapResLike, err := biz.restaurantLikestore.GetRestaurantLikes(ctx, ids)
	if err != nil {
		log.Println("Cannot get restaurant likes:", err)
	}
	if v := mapResLike; v != nil {
		for i, item := range result {
			result[i].LikedCount = mapResLike[item.Id]
		}
	}
	return result, nil
}
