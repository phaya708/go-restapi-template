package repository

import (
	"go-restapi-template/app/domain/entity"
	"go-restapi-template/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetByID_Valid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)

	id := uint(1)
	userEntity := entity.User{
		ID: &id,
	}
	
	user, err := userRepository.GetByID(userEntity)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, validUserData[0].FirstName, user.FirstName)
	assert.Equal(t, validUserData[0].LastName, user.LastName)
}

func TestUserRepository_GetByID_Invalid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)

	id := uint(10)
	userEntity := entity.User{
		ID: &id,
	}
	
	user, err := userRepository.GetByID(userEntity)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserRepository_GetAll(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)
	users, err := userRepository.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 3, len(*users))

	for i, user := range *users {
		assert.Equal(t, uint(i+1), user.ID)
		assert.Equal(t, user.FirstName, validUserData[i].FirstName)
		assert.Equal(t, user.LastName, validUserData[i].LastName)
	}
}

func TestUserRepository_Create(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)

	firstName := "test"
	lastName := "test"

	userEntity := entity.User{
		FirstName: &firstName,
		LastName: &lastName,
	}

	user, err := userRepository.Create(userEntity)

	assert.NoError(t, err)
	assert.Equal(t, uint(4), user.ID)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, lastName, user.LastName)
}

func TestUserRepository_Update_Valid_AllParams(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)

	id := uint(1)
	firstName := "test"
	lastName := "test"

	userEntity := entity.User{
		ID: &id,
		FirstName: &firstName,
		LastName: &lastName,
	}

	user, err := userRepository.Update(userEntity)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, lastName, user.LastName)
}

func TestUserRepository_Update_Valid_OnlyFirstName(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)

	id := uint(1)
	firstName := "test"

	userEntity := entity.User{
		ID: &id,
		FirstName: &firstName,
	}

	user, err := userRepository.Update(userEntity)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, "太郎", user.LastName)
}

func TestUserRepository_Delete_Valid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)

	id := uint(1)
	userEntity := entity.User{
		ID: &id,
	}

	err := userRepository.Delete(userEntity)

	assert.NoError(t, err)
}

func TestUserRepository_Delete_Invalid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := NewUserRepository(db)

	id := uint(20)
	userEntity := entity.User{
		ID: &id,
	}

	err := userRepository.Delete(userEntity)

	assert.Error(t, err)
}