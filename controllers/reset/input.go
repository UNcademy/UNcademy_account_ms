package reset

type InputReset struct {
	Password    string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

//Cambiar la contraseña
