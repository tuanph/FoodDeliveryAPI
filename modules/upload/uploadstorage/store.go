package uploadstorage

import "gorm.io/gorm"

type mySqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *mySqlStore {
	return &mySqlStore{db: db}
}
