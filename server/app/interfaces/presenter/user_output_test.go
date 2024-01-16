package presenter

import (
	"go-restapi-template/app/domain/entity"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUserOutput(t *testing.T) {
	id := uint(1)
	firstName := "test"
	lastName := "test"

	userEntity := entity.User{
		ID: &id,
		FirstName: &firstName,
		LastName: &lastName,
	}

	output := NewUserOutput(userEntity)

	assert.Equal(t, id, output.ID)
	assert.Equal(t, firstName, output.FirstName)
	assert.Equal(t, lastName, output.LastName)
}

func TestUsersOutput(t *testing.T) {

	var usersEntity entity.Users

	for i := 0; i < 3; i++ {
		id := uint(i)
		firstName := "test"
		lastName := "test"

		userEntity := entity.User{
			ID: &id,
			FirstName: &firstName,
			LastName: &lastName,
		}

		usersEntity = append(usersEntity, userEntity)
	}

	output := NewUsersOutput(usersEntity)

	assert.Equal(t, 3, len(output))

	for i, user := range usersEntity {
		assert.Equal(t, *user.ID, output[i].ID)
		assert.Equal(t, *user.FirstName, output[i].FirstName)
		assert.Equal(t, *user.LastName, output[i].LastName)
	}
}