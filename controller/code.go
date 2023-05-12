package controller

import (
	"code_site/service"

	"github.com/gin-gonic/gin"
)

func ControlOther(group *gin.RouterGroup) {
	//	service.GetCodeHandler
	group.GET("/captcha", service.GetCodeServer)
}
