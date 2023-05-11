package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCodeHandler(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, map[string]string{
		"code": "2000",
	})
}
