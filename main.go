package main

import (
	"github.com/DarlanPrado-Access/primeiro-projeto-go/config"
	"github.com/DarlanPrado-Access/primeiro-projeto-go/router"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Erro de inicialização de configuração: %v", err)
		return
	}

	gin.SetMode(gin.TestMode)

	router.Initialize()
}
