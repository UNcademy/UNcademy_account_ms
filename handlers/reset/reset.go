package reset

import (
	reset2 "UNcademy_account_ms/controllers/reset"
	util "UNcademy_account_ms/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type handler struct {
	service reset2.Service
}

func NewHandlerReset(service reset2.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResetHandler(ctx *gin.Context) {
	var input reset2.InputReset

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		defer logrus.Error(err.Error())
		util.APIResponse(ctx, "Parsing json data failed", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		resultToken, errToken := util.VerifyTokenHeader(ctx, "JWT_SECRET")

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Verified activation token failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		result := util.DecodeToken(resultToken)

		errReset := h.service.ResetService(&input, result.Claims.UserName)

		switch errReset {

		case "USER_NOT_FOUND_404":
			util.APIResponse(ctx, "User not found", http.StatusNotFound, http.MethodPut, nil)
		case "USER_NOT_ACTIVE_403":
			util.APIResponse(ctx, "User not active", http.StatusForbidden, http.MethodPut, nil)
		case "WRONG_PASSWORD_403":
			util.APIResponse(ctx, "Incorrect password", http.StatusForbidden, http.MethodPut, nil)
		case "CHANGE_PASSWORD_FAILED_403":
			util.APIResponse(ctx, "Change password failed", http.StatusForbidden, http.MethodPut, nil)
		default:
			util.APIResponse(ctx, "Password changed successfully!", http.StatusOK, http.MethodPut, nil)
		}
	}
}
