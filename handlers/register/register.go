package register

import (
	register2 "UNcademy_account_ms/controllers/register"
	util "UNcademy_account_ms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type handler struct {
	service  register2.Service
	rabbitmq *amqp.Channel
}

func NewHandlerRegister(service register2.Service, rabbitmq *amqp.Channel) *handler {
	return &handler{service: service, rabbitmq: rabbitmq}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	//Extrae el json
	var input register2.InputRegister
	ctx.ShouldBindJSON(&input)

	resultRegister, errRegister := h.service.RegisterService(&input)

	switch errRegister {

	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:

		util.SendMessage(h.rabbitmq, resultRegister.UserName, resultRegister.Program)
		util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
