package utils

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// llega contraseña encriptada y debe ser igual a el hash de la que llego
func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

// Cuando hago el registro, se llama a esta funcion para guardar la contraseña en la bd encriptada
func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return string(result)
}

//hash: convierte en codigo lo que entre (en este caso la contraseña) y no es invertible
