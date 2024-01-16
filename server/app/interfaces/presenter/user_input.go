package presenter

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type GetUserInput struct {
	ID uint
}

type CreateUserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}

type UpdateUserInput struct {
	ID uint `binding:"-"`
	FirstName *string `json:"first_name" binding:"-"`
	LastName *string `json:"last_name" binding:"-"`
}

type DeleteUserInput struct {
	ID uint `json:"id" binding:"required"`
}

func NewCreateUserInput(c *gin.Context) (*CreateUserInput, error) {
	var inputData CreateUserInput
	if err := c.BindJSON(&inputData); err != nil {
		return nil, err
	}
	return &inputData, nil
}

func NewUpdateUserInput(c *gin.Context) (*UpdateUserInput, error) {
	var inputData UpdateUserInput
	if err := c.BindJSON(&inputData); err != nil {
		return nil, err
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return nil, err
	}
	inputData.ID = uint(id)

	return &inputData, nil
}

func NewGetUserInput(c *gin.Context) (*GetUserInput, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return nil, err
	}

	inputData := GetUserInput{
		ID: uint(id),
	}

	return &inputData, nil
}

func NewDeleteUserInput(c *gin.Context) (*DeleteUserInput, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		return nil, err
	}

	inputData := DeleteUserInput{
		ID: uint(id),
	}

	return &inputData, nil
}
