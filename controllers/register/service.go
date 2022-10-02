package register

import model "UNcademy_account_ms/models"

type Service interface {
	RegisterService(input *InputRegister) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*model.User, string) {

	users := model.User{
		UserName:       input.UserName,
		UserType:       input.UserType,
		Password:       input.Password,
		FullName:       input.FullName,
		Email:          input.Email,
		Document:       input.Document,
		DepDocument:    input.DepDocument,
		CityDocument:   input.CityDocument,
		Genre:          input.Genre,
		UNMail:         input.UNMail,
		Cel:            input.Cel,
		Tel:            input.Tel,
		Age:            input.Age,
		BirthPlace:     input.BirthPlace,
		Country:        input.Country,
		BloodType:      input.BloodType,
		Address:        input.Address,
		ArmyCard:       input.ArmyCard,
		MotherFullName: input.MotherFullName,
		MotherDocument: input.MotherDocument,
		FatherFullName: input.FatherFullName,
		FatherDocument: input.FatherDocument,
		Program:        input.Program,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}
