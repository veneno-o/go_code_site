package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func noMatchHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
}
