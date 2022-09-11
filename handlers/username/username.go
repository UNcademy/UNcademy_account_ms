package username

import (
	util "UNcademy_account_ms/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct{}

func NewHandlerAddAddress() *handler {
	return &handler{}
}

func (h *handler) GetUsernameHandler(ctx *gin.Context) {
	resultToken, errToken := util.VerifyTokenHeader(ctx, "JWT_SECRET")

	if errToken != nil {
		defer logrus.Error(errToken.Error())
		util.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodGet, nil)
		return
	}

	result := util.DecodeToken(resultToken)

	util.APIResponse(ctx, "Valid token returned username", http.StatusOK, http.MethodGet, result.Claims.UserName)
}

//Endpoint: recibe el jwt (es un token) y devuelve el nombre de usuario asociado a ese token
//jwt se utiliza para la generacion de tokens de acceso
