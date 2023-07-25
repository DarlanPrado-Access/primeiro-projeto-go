package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExempleController struct{}

func (c *ExempleController) GetTeste(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "mensagem exemplo",
	})
}
