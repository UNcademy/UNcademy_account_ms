package reset

import (
	model "UNcademy_account_ms/models"
	util "UNcademy_account_ms/utils"
	"gorm.io/gorm"
)

type Repository interface {
	ResetRepository(input *model.User, newPassword string) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryReset(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResetRepository(input *model.User, newPassword string) (*model.User, string) {
	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.UserName = input.UserName
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("user_name = ?", input.UserName).Find(&users)

	if checkUserAccount.RowsAffected == 0 {
		errorCode <- "USER_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "USER_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "WRONG_PASSWORD_403"
		return &users, <-errorCode
	}

	changeUserPassword := db.Debug().Select("password").Where("user_name = ?", input.UserName).Update(
		"password", util.HashPassword(newPassword))

	if changeUserPassword.Error != nil {
		errorCode <- "CHANGE_PASSWORD_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
