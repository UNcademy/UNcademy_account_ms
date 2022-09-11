package validation

import (
	util "UNcademy_account_ms/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateHandler(ctx *gin.Context) {
	util.APIResponse(ctx, "Token validated!", http.StatusOK, http.MethodGet, nil)
}

//Endpoint que envia la respuesta de que el token esta bien
//endpoint (es el punto de acceso a la app => una puerta de una casa y cada puerta lleva a un cuarto especifico)
