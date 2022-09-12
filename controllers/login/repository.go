package login

import (
	model "UNcademy_account_ms/models"
	util "UNcademy_account_ms/utils"
	"gorm.io/gorm"
)

type Repository interface {
	LoginRepository(input *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

// Todos los repositorios tienen relacion directa con la base de datos
func (r *repository) LoginRepository(input *model.User) (*model.User, string) {
	//Es el modelo de usuario que tenemos definidos pero nulos (estructura temporal)
	var users model.User
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.UserName = input.UserName
	users.Password = input.Password

	//Miro si el nombre de usuario existe en la base de datos y guarda en users (estructura temporal) la informacion que esta guardada de ese usuario
	checkUserAccount := db.Debug().Select("*").Where("user_name = ?", input.UserName).Find(&users)

	if checkUserAccount.RowsAffected == 0 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	//Compara la contraseña que llego con la que esta guardada en la base de datos
	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
