package register

import (
	register2 "UNcademy_account_ms/controllers/register"
	util "UNcademy_account_ms/utils"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"net/http"
)

type handler struct {
	service  register2.Service
	rabbitmq *amqp.Channel
}

func NewHandlerRegister(service register2.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	//Extrae el json
	var input register2.InputRegister
	ctx.ShouldBindJSON(&input)

	_, errRegister := h.service.RegisterService(&input)

	switch errRegister {

	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:

		util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
