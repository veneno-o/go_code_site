package main

import (
	"code_site/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	dao.Connect()
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": 123,
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/", index)
	engine.Run(":7788")
}
