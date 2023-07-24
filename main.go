package main

import (
	"github.com/DarlanPrado-Access/primeiro-projeto-go/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.TestMode)

	router.Initialize()
}
