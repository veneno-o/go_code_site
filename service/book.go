package service

import (
	"code_site/dao"
	"code_site/model"
	"code_site/utils"
	"strconv"
)

type Res struct {
	Data   []*model.Book `json:"data"`
	ErrMsg string        `json:"errMsg"`
}

// AddBook 添加书籍
func AddBook(book *model.Book) Res {
	// 1. 数据验证
	res := verify(book)
	if res != "" {
		return Res{nil, res}
	}
	//2.验证通过，向数据库插入数据
	err := dao.AddBook(book)
	if err != nil {
		return Res{nil, err.Error()}
	} else {
		return Res{nil, ""}
	}
}

// GetBookById DelBook 根据id删除书籍
// func DelBook(id int) res {
//
// }
//
// // UpdateBookById 根据id修改书籍
// func UpdateBookById(id int, book *dao.Book) res {
//
// }
//
// GetBookById 根据id获取书籍
func GetBookById(id string) Res {
	// 1.id类型转换
	newId, err := strconv.Atoi(id)

	if err != nil {
		return Res{nil, "book id 类型有误"}
	}

	//2.向持久层获取数据
	var bookSlice []*model.Book
	book, err := dao.GetBookById(newId)

	if err != nil {
		return Res{nil, err.Error()}
	} else {
		bookSlice = append(bookSlice, book)
		return Res{bookSlice, ""}
	}
}

// // GetBookByPage 分页获取书籍
// func GetBookByPage() res {
//
// }
func verify(book *model.Book) string {
	if book.BookTitle == "" {
		return "书籍标题格式有误"
	}
	if book.BookPic == "" {
		return "书籍图片格式有误"
	}
	if book.DownloadLink == "" {
		return "下载链接格式有误"
	}
	if book.BookIntro == "" {
		return "书籍介绍格式有误"
	}
	if book.OnShelfDate == "" {
		return "书籍上架日期格式有误"
	}
	if book.ScanNumber <= 0 {
		return "书籍浏览数必须为正整数"
	}
	if book.CommentNumber <= 0 {
		return "书籍评论数必须为正整数"
	}
	if book.RequirePoints <= 0 {
		return "书籍下载所需积分必须为正整数"
	}
	if book.TypeId == 0 {
		return "书籍所属分类id不能为0"
	}
	//判断上架日期是否有效
	timeStr := utils.GetCurTime()
	if book.OnShelfDate > timeStr {
		return "书籍上架日期有误"
	}
	//所属分类id是否有效
	return ""
}
