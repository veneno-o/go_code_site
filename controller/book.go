package controller

import (
	"code_site/model"
	"code_site/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type res struct {
	code   int
	ErrMsg string // 错误信息
	Data   []model.Book
}

func BookController(group *gin.RouterGroup) {

	//	添加书籍
	group.POST("/", AddBook)

	//	根据id删除书籍
	group.DELETE("/:id", DelBook)

	// 根据id修改书籍
	group.PATCH("/:id", DelBook)

	//	根据id获取书籍
	group.GET("/:id", GetBookById)

	//	分页获取书籍
	group.GET("/", GetBookByPage)

}

// AddBook 添加书籍
func AddBook(ctx *gin.Context) {
	var book model.Book
	ctx.BindJSON(&book)
	res := service.AddBook(&book)
	if res.ErrMsg != "" {
		ctx.JSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

// DelBook 根据id删除书籍
func DelBook(ctx *gin.Context) {

}

// UpdateBookById 根据id修改书籍
func UpdateBookById(ctx *gin.Context) {

}

// GetBookById 根据id获取书籍
func GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	res := service.GetBookById(id)
	if res.ErrMsg != "" {
		ctx.JSON(http.StatusNotFound, res)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

// GetBookByPage 分页获取书籍
func GetBookByPage(ctx *gin.Context) {

}
