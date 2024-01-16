package presenter

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-restapi-template/testutil"
)

func TestUserInput_Get_Valid(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	request, err := NewGetUserInput(c)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), request.ID)
}

func TestUserInput_Create_Valid(t *testing.T) {
	requestBody, err := json.Marshal(testutil.UserRequestBody{
		FirstName: "test",
		LastName: "test",
	})
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	
	request, err := NewCreateUserInput(c)
	assert.Nil(t, err)
	assert.Equal(t, "test", request.FirstName)
	assert.Equal(t, "test", request.LastName)
}

func TestUserInput_Create_Invalid(t *testing.T) {	
	requestBody, err := json.Marshal(testutil.UserRequestBody{
		FirstName: "",
		LastName: "",
	})
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	request, err := NewCreateUserInput(c)

	assert.Error(t, err)
	assert.Nil(t, request)
}

func TestUserInput_Update_AllParams(t *testing.T) {
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

	request, err := NewUpdateUserInput(c)

	assert.NoError(t, err)
	assert.Equal(t, "test", *request.FirstName)
	assert.Equal(t, "test", *request.LastName)
}

func TestUserInput_Update_OnlyFirstName(t *testing.T) {
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

	request, err := NewUpdateUserInput(c)

	assert.NoError(t, err)
	assert.Equal(t, "test", *request.FirstName)
	assert.Nil(t, request.LastName)
}

func TestUserInput_Delete_Valid(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	request, err := NewDeleteUserInput(c)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), request.ID)
}