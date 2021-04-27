package ginrestaurant

import (
	"FoodDelivery/common"
	"FoodDelivery/component"
	"FoodDelivery/modules/restaurant/restaurantbiz"
	"FoodDelivery/modules/restaurant/restaurantmodel"
	"FoodDelivery/modules/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewMySQLStore(appCtx.GetConnectionString())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}
		for i := range result {
			result[i].Mask(false)
			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
