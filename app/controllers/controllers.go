package controllers

import "github.com/gin-gonic/gin"

//lista todos os controllers e quais funções de cada controller podem ter acesso externo

type ExempleControllerInterface interface {
	GetTeste(ctx *gin.Context)
}
