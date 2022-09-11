package register

import (
	model "UNcademy_account_ms/models"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.User) (*model.User, string) {

	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	//Estoy verificando que el usuario no haya sido registrado antes
	checkUserAccount := db.Debug().Select("*").Where("user_name = ?", input.UserName).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.UserName = input.UserName
	users.UserType = input.UserType
	users.Password = input.Password
	users.FullName = input.FullName
	users.Document = input.Document
	users.DepDocument = input.DepDocument
	users.CityDocument = input.CityDocument
	users.Genre = input.Genre
	users.Email = input.Email
	users.UNMail = input.UNMail
	users.Cel = input.Cel
	users.Tel = input.Tel
	users.Age = input.Age
	users.BirthPlace = input.BirthPlace
	users.Country = input.Country
	users.BloodType = input.BloodType
	users.Address = input.Address
	users.ArmyCard = input.ArmyCard
	users.MotherFullName = input.MotherFullName
	users.MotherDocument = input.MotherDocument
	users.FatherFullName = input.FatherFullName
	users.FatherDocument = input.FatherDocument

	//Funciones de ORM
	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode

}
