package presenter

import (
	"go-restapi-template/app/domain/entity"
)

type UserOutput struct {
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type UsersOutput []UserOutput

func NewUserOutput(user entity.User) (UserOutput) {
	output := UserOutput{
		ID: *user.ID,
		FirstName: *user.FirstName,
		LastName: *user.LastName,
	}
	return output
}

func NewUsersOutput(users entity.Users) (UsersOutput) {
	var output UsersOutput
	for _, user := range users {
		output = append(output, UserOutput{
			ID: *user.ID,
			FirstName: *user.FirstName,
			LastName: *user.LastName,
		})
	}
	return output
}