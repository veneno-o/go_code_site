package dao

import (
	"code_site/model"
	"log"
)

// AddBook 添加书籍
func AddBook(book *model.Book) error {
	result := Db.Create(&book)
	if result.Error != nil {
		return result.Error
	} else {
		log.Printf("AddBook:err>>>%t\n", result.Error)
		return nil
	}
}

// DelBook 根据id删除书籍
// func DelBook(id int) res {
//
// }
//
// // UpdateBookById 根据id修改书籍
// func UpdateBookById(id int, book *dao.Book) res {
//
// }

// GetBookById 根据id获取书籍
func GetBookById(id int) (*model.Book, error) {
	var book *model.Book
	tx := Db.Debug().First(&book, id)

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return nil, tx.Error
	} else {
		return book, nil
	}
}

// // GetBookByPage 分页获取书籍
// func GetBookByPage() res {
//
// }
