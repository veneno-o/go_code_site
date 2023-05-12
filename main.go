package main

import (
	"code_site/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	controller.Controller(engine)
	engine.Run(":7788")
}
