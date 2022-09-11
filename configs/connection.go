package configs

import (
	model "UNcademy_account_ms/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Connection() *gorm.DB {
	//Datos para hacer la conexión con la base de datos y devuelve un objeto de tipo gorm.DB
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	name := os.Getenv("DB_NAME")

	urn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	db, err := gorm.Open(mysql.Open(urn), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
