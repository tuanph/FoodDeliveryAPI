package component

import (
	"FoodDelivery/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetConnectionString() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}
func (ctx *appCtx) GetConnectionString() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
