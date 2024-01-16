package usecase

import (
	"go-restapi-template/app/domain/entity"
	"go-restapi-template/app/interfaces/presenter"
	"go-restapi-template/app/interfaces/repository"
)

type UserUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return UserUsecase{ur}
}

func (uu *UserUsecase) GetAll() (*presenter.UsersOutput, error) {
	usersData, err := uu.ur.GetAll()
	
	if err != nil {
		return nil, err
	}

	var usersEntity entity.Users

	for _, userData := range *usersData {
		id := userData.ID
    firstName := userData.FirstName
    lastName := userData.LastName

    user := entity.User{
        ID: &id,
        FirstName: &firstName,
        LastName: &lastName,
    }
		usersEntity = append(usersEntity, user)
	}

	outputData := presenter.NewUsersOutput(usersEntity)

	return &outputData, nil
}

func (uu *UserUsecase) GetByID(inputData presenter.GetUserInput) (*presenter.UserOutput, error) {
	userEntity := entity.User{
		ID: &inputData.ID,
	}

	userData, err := uu.ur.GetByID(userEntity)
	if err != nil {
		return nil, err
	}

	usersEntity := entity.User{
		ID: &userData.ID,
		FirstName: &userData.FirstName,
		LastName: &userData.LastName,
	}

	outputData := presenter.NewUserOutput(usersEntity)	
	return &outputData, nil
}

func (uu *UserUsecase) Create(inputData presenter.CreateUserInput) (*presenter.UserOutput, error) {
	user := entity.User{
		FirstName: &inputData.FirstName,
		LastName: &inputData.LastName,
	}

	userData, err := uu.ur.Create(user)
	if err != nil {
		return nil, err
	}

	outputData := presenter.UserOutput{
		ID: userData.ID,
		FirstName: userData.FirstName,
		LastName: userData.LastName,
	}

	return &outputData, nil
}

func (uu *UserUsecase) Update(inputData presenter.UpdateUserInput) (*presenter.UserOutput, error) {
	user := entity.User{
		ID: &inputData.ID,
		FirstName: inputData.FirstName,
		LastName: inputData.LastName,
	}

	userData, err := uu.ur.Update(user)
	if err != nil {
		return nil, err
	}

	outputData := presenter.UserOutput{
		ID: userData.ID,
		FirstName: userData.FirstName,
		LastName: userData.LastName,
	}

	return &outputData, nil
}

func (uu *UserUsecase) Delete(inputData presenter.DeleteUserInput) error {
	userEntity := entity.User{
		ID: &inputData.ID,
	}
	err := uu.ur.Delete(userEntity)
	if err != nil {
		return err
	}
	return nil
}