package main

import (
	"code_site/controller"
	"code_site/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	dao.Connect()
}

func main() {
	engine := gin.Default()
	controller.RouterHandler(engine)
	engine.Run(":7788")
}
