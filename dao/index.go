package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const username, password, host, port, Dbname = "root", "123456", "127.0.0.1", 3306, "code_site"

var Db *gorm.DB
var Err error

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	Db, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err != nil {
		log.Println("数据库连接失败")
		panic(Err)
	} else {
		log.Println("mysql数据库: code_site 连接成功")
	}
	//2.迁移表结构
	Db.Migrator().AutoMigrate(&Admin{}, &Book{}, &Comment{}, &Interview{}, &Issue{}, &Types{}, &User{})
}
