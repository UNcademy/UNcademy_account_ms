package reset

import model "UNcademy_account_ms/models"

type Service interface {
	ResetService(input *InputReset, username string) string
}

type service struct {
	repository Repository
}

func NewServiceReset(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResetService(input *InputReset, username string) string {
	user := model.User{
		UserName: username,
		Password: input.Password,
	}

	_, errReset := s.repository.ResetRepository(&user, input.NewPassword)

	return errReset
}
