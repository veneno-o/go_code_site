package controller

import (
	"code_site/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var BookRouterGroup *gin.RouterGroup
var CommentRouterGroup *gin.RouterGroup
var InterviewRouterGroup *gin.RouterGroup
var IssueRouterGroup *gin.RouterGroup
var TypeRouterGroup *gin.RouterGroup
var UserRouterGroup *gin.RouterGroup

func noMatchHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
}

func RouterHandler(ginServe *gin.Engine) {
	//root, _ := os.Getwd()

	//ginServe.Static("/css", root+"/static/css")
	//ginServe.Static("/fonts", root+"/static/fonts")
	//ginServe.Static("/js", root+"/static/js")

	ginServe.LoadHTMLGlob("templates/*")

	// response index page
	ginServe.GET("/", service.GetPageHandler)

	// response data
	routerGroup := ginServe.Group("/api")
	{
		BookRouterGroup = routerGroup.Group("/book")
		CommentRouterGroup = routerGroup.Group("/comment")
		CommentRouterGroup = routerGroup.Group("/interview")
		CommentRouterGroup = routerGroup.Group("/issue")
		CommentRouterGroup = routerGroup.Group("/type")
		CommentRouterGroup = routerGroup.Group("/user")
	}

	//get code获取验证码
	ginServe.GET("/res/captcha", service.GetCodeHandler)

	// no match response 404.html
	ginServe.NoRoute(noMatchHandler)
}
