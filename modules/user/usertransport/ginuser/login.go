package ginuser

import (
	"FoodDelivery/common"
	"FoodDelivery/component"
	"FoodDelivery/component/hasher"
	"FoodDelivery/component/tokenprovider/jwt"
	"FoodDelivery/modules/user/userbiz"
	"FoodDelivery/modules/user/usermodel"
	"FoodDelivery/modules/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetConnectionString()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewMySQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
