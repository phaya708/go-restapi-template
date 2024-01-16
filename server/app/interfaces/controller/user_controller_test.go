package controller

import (
	"go-restapi-template/app/domain/usecase"
	"go-restapi-template/app/interfaces/repository"
	"go-restapi-template/app/interfaces/presenter"
	"encoding/json"
	"bytes"
	"go-restapi-template/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserController_GetByID_ValidValue(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.GetByID(c)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := presenter.UserOutput{
		ID: validUserData[0].ID,
		FirstName: validUserData[0].FirstName,
		LastName: validUserData[0].LastName,
	}

	var actualResponse presenter.UserOutput
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestUserController_GetByID_InvalidValue(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "abc"}}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.GetByID(c)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserController_GetAll(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()

	testutil.MigrateUser(db, validUserData)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.GetAll(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse presenter.UsersOutput
	for _, user := range validUserData {
		expectedResponse = append(expectedResponse, presenter.UserOutput{
			ID: user.ID,
			FirstName: user.FirstName,
			LastName: user.LastName,
		})
	}

	var actualResponse presenter.UsersOutput
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestUserController_Create_Valid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	requestBody, err := json.Marshal(testutil.UserRequestBody{
		FirstName: "test",
		LastName: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.Create(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUserController_Create_Invalid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	requestBody, err := json.Marshal(testutil.UserRequestBody{
		FirstName: "",
		LastName: "",
	})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.Create(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserController_Update_Valid_AllParam(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	requestBody, err := json.Marshal(testutil.UserRequestBody{
		FirstName: "test",
		LastName: "test",
	})

	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.Update(c)

	assert.Equal(t, http.StatusOK, w.Code)
}


func TestUserController_Update_Valid_OneParam(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	requestBody, err := json.Marshal(testutil.UserRequestBody{
		FirstName: "test",
	})
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.Update(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_Update_Invalid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)

	requestBody, err := json.Marshal(testutil.UserRequestBody{
		FirstName: "test",
		LastName: "test",
	})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/users/abc", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "abc"}}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.Update(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserController_Delete_Valid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)
	
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.Delete(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_Delete_Invalid(t *testing.T) {
	db := testutil.NewTestDB()
	validUserData := testutil.ValidUserData()
	testutil.MigrateUser(db, validUserData)
	
	req, err := http.NewRequest("DELETE", "/users/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "abc"}}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := NewUserController(userUsecase)

	userController.Delete(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}