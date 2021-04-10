package component

import "gorm.io/gorm"

type AppContext interface {
	GetConnectionString() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}
func (ctx *appCtx) GetConnectionString() *gorm.DB {
	return ctx.db
}
