package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicia as rotas e suas pré, configurações

	router := gin.Default()

	InitializeRouter(router)

	router.Run()
}
