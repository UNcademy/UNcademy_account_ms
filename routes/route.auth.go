package route

import (
	login2 "UNcademy_account_ms/controllers/login"
	register2 "UNcademy_account_ms/controllers/register"
	reset2 "UNcademy_account_ms/controllers/reset"
	handlerLogin "UNcademy_account_ms/handlers/login"
	handlerRegister "UNcademy_account_ms/handlers/register"
	handlerReset "UNcademy_account_ms/handlers/reset"
	handlerGetUsername "UNcademy_account_ms/handlers/username"
	validationHandler "UNcademy_account_ms/handlers/validation"
	"UNcademy_account_ms/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	loginRepository := login2.NewRepositoryLogin(db)
	loginService := login2.NewServiceLogin(loginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := register2.NewRepositoryRegister(db)
	registerService := register2.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	resetRepository := reset2.NewRepositoryReset(db)
	resetService := reset2.NewServiceReset(resetRepository)
	resetHandler := handlerReset.NewHandlerReset(resetService)

	getUsernameHandler := handlerGetUsername.NewHandlerAddAddress()

	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.PUT("/reset", resetHandler.ResetHandler)
	groupRoute.GET("/validate", middlewares.Auth(), validationHandler.ValidateHandler)
	groupRoute.GET("/username", getUsernameHandler.GetUsernameHandler)
}

//Post: Crear - o cuando el request llega en un json
//Get: read
//Put: Update

//Flujo general:
//Request (route)
//Handle
//Service
//Repository
//Variable temporal
