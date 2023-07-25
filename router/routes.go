package router

import (
	"github.com/DarlanPrado-Access/primeiro-projeto-go/app/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {

	//declara os controllers onde há funções para as rotas acessarem
	exempleController := &controllers.ExempleController{}

	//determina as rotas de acesso da api
	api := router.Group("/api")
	{
		api.GET("/teste", exempleController.GetTeste)
	}
}
