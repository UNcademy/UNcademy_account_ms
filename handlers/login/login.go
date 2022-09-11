package login

import (
	login2 "UNcademy_account_ms/controllers/login"
	util "UNcademy_account_ms/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	service login2.Service
}

func NewHandlerLogin(service login2.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	//Define username y password de los usuarios
	var input login2.InputLogin

	//Json que me llega en el request (dentro del request esta el json)
	err := ctx.ShouldBindJSON(&input)

	//Error extrayendo el json (esta mal)
	if err != nil {
		defer logrus.Error(err.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {

		resultLogin, errLogin := h.service.LoginService(&input)

		switch errLogin {

		case "LOGIN_NOT_FOUND_404":
			util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)

		case "LOGIN_NOT_ACTIVE_403":
			util.APIResponse(ctx, "User account is not active", http.StatusForbidden, http.MethodPost, nil)

		case "LOGIN_WRONG_PASSWORD_403":
			util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)

		default:
			accessTokenData := map[string]interface{}{"username": resultLogin.UserName, "email": resultLogin.Email, "usertype": resultLogin.UserType}
			accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 300)

			if errToken != nil {
				defer logrus.Error(errToken.Error())
				util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			}

			util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": accessToken})
		}
	}
}
