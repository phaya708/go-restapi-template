package reset

import (
	"gorm.io/gorm"
	"go-restapi-template/db/model"
)

func Reset(db *gorm.DB) {
	db.Migrator().DropTable(&model.User{})
	db.Migrator().AutoMigrate(&model.User{})
}