package controller

import (
	"code_site/service"
	"github.com/gin-gonic/gin"
)

func Controller(ginServe *gin.Engine) {
	//root, _ := os.Getwd()

	//ginServe.Static("/css", root+"/static/css")
	//ginServe.Static("/fonts", root+"/static/fonts")
	//ginServe.Static("/js", root+"/static/js")

	ginServe.LoadHTMLGlob("templates/*")

	// request index page
	ginServe.GET("/", service.GetPageHandler)

	// request database data
	routerGroup := ginServe.Group("/api")
	{
		// control book
		BookGroup := routerGroup.Group("/book")
		BookController(BookGroup)

		// control comment
		CommentGroup := routerGroup.Group("/comment")
		CommentController(CommentGroup)

		// control interview
		InterviewGroup := routerGroup.Group("/interview")
		InterviewController(InterviewGroup)

		// control issue
		IssueGroup := routerGroup.Group("/issue")
		IssueController(IssueGroup)

		// control type
		TypeGroup := routerGroup.Group("/type")
		TypeController(TypeGroup)

		//control user
		UserGroup := routerGroup.Group("/user")
		UserController(UserGroup)
	}

	// request other data
	otherGroup := ginServe.Group("/res")
	ControlOther(otherGroup)

	// no match response 404.html
	ginServe.NoRoute(noMatchHandler)
}
