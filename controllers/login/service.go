package login

import (
	model "UNcademy_account_ms/models"
)

type Service interface {
	LoginService(input *InputLogin) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

// Transforma el request en el modelo que tengo definido en mi base de datos
func (s *service) LoginService(input *InputLogin) (*model.User, string) {
	user := model.User{
		UserName: input.UserName,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
