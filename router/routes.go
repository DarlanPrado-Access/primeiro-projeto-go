package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/teste", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "testado",
			})
		})
	}
}
