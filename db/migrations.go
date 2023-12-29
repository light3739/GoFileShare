package db

import (
	"github.com/jinzhu/gorm"
	"share-files/data"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&data.File{})
}
