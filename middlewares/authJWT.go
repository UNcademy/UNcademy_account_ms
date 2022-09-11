package middlewares

import (
	util "UNcademy_account_ms/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UnAuthorizeError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var errorResponse UnAuthorizeError

		errorResponse.Status = "Forbidden"
		errorResponse.Code = http.StatusForbidden
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "Authorization is required for this endpoint"

		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusForbidden, errorResponse)
			defer ctx.AbortWithStatus(http.StatusForbidden)
		}

		token, err := util.VerifyTokenHeader(ctx, "JWT_SECRET")

		errorResponse.Status = "UnAuthorize"
		errorResponse.Code = http.StatusUnauthorized
		errorResponse.Method = ctx.Request.Method
		errorResponse.Message = "accessToken invalid or expired"

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			ctx.Set("user", token.Claims)
			ctx.Next()
		}
	}
}

//middleware: es el primer paso de la entrada entre el endpoint y el handler
//Realmente valida que el token este correcto para que la cuando se envie la respuesta : ok
