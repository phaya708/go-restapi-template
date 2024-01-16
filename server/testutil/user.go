package testutil

import (
	"go-restapi-template/db/model"
	"gorm.io/gorm"
)

type UserRequestBody struct {
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
}

func ValidUserData() []model.User {
	users := model.Users{
		{FirstName: "山田", LastName: "太郎"},
		{FirstName: "佐藤", LastName: "次郎"},
		{FirstName: "田中", LastName: "三郎"},
	}
	return users
}

func MigrateUser(db *gorm.DB, users []model.User) {
	db.Migrator().DropTable(&model.User{})
	db.AutoMigrate(&model.User{})
	db.Create(&users)
}


