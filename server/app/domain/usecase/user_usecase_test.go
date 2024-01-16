package usecase

import (
	"go-restapi-template/app/interfaces/presenter"
	"go-restapi-template/app/interfaces/repository"
	"go-restapi-template/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_GetByID_Valid(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepository)

	inputData := presenter.GetUserInput{
		ID: uint(1),
	}

	user, err := userUsecase.GetByID(inputData)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, validUserData[0].FirstName, user.FirstName)
	assert.Equal(t, validUserData[0].LastName, user.LastName)

}

func TestUserUsecase_GetByID_Invalid(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepository)

	inputData := presenter.GetUserInput{
		ID: uint(10),
	}

	user, err := userUsecase.GetByID(inputData)

	assert.Error(t, err)
	assert.Nil(t, user)
}


func TestUserUsecase_GetAll(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepository)

	users, err := userUsecase.GetAll()

	assert.NoError(t, err)

	for i, user := range *users {
		assert.Equal(t, uint(i+1), user.ID)
		assert.Equal(t, user.FirstName, validUserData[i].FirstName)
		assert.Equal(t, user.LastName, validUserData[i].LastName)
	}
}

func TestUserUsecase_Create(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepository)

	inputData := presenter.CreateUserInput{
		FirstName: "test",
		LastName: "test",
	}

	user, err := userUsecase.Create(inputData)

	assert.NoError(t, err)
	assert.Equal(t, uint(4), user.ID)
	assert.Equal(t, inputData.FirstName, user.FirstName)
	assert.Equal(t, inputData.LastName, user.LastName)
}

func TestUserUsecase_Update_Valid_AllParams(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepository)

	firstName := "test"
	lastName := "test"

	inputData := presenter.UpdateUserInput{
		ID: uint(1),
		FirstName: &firstName,
		LastName: &lastName,
	}

	user, err := userUsecase.Update(inputData)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, lastName, user.LastName)
}

func TestUserUsecase_Update_Valid_OnlyFirstName(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepository)

	firstName := "test"

	inputData := presenter.UpdateUserInput{
		ID: uint(1),
		FirstName: &firstName,
	}

	user, err := userUsecase.Update(inputData)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, validUserData[0].LastName, user.LastName)
}

func TestUserUsecase_Delete_Valid(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepository)

	inputData := presenter.DeleteUserInput{
		ID: uint(1),
	}

	err := userUsecase.Delete(inputData)

	assert.NoError(t, err)
}

func TestUserUsecase_Delete_Invalid(t *testing.T){
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	userRepository := repository.NewUserRepository(db)

	inputData := presenter.DeleteUserInput{
		ID: uint(10),
	}
	userUsecase := NewUserUsecase(userRepository)

	err := userUsecase.Delete(inputData)

	assert.Error(t, err)
}
