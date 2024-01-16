package controller

import (
	"go-restapi-template/app/interfaces/presenter"
	"go-restapi-template/app/domain/usecase"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type UserController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu}
}

func (uc *userController) GetAll(c *gin.Context) {
	response, err := uc.uu.GetAll()

	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (uc *userController) GetByID(c *gin.Context) {
	inputData, err := presenter.NewGetUserInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := uc.uu.GetByID(*inputData)
	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (uc *userController) Create(c *gin.Context) {
	request, err := presenter.NewCreateUserInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := uc.uu.Create(*request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *userController) Update(c *gin.Context) {
	inputData, err := presenter.NewUpdateUserInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := uc.uu.Update(*inputData)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (uc *userController) Delete(c *gin.Context) {
	inputData, err := presenter.NewDeleteUserInput(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = uc.uu.Delete(*inputData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}