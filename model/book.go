package model

type Book struct {
	Id            int    `json:"id"`
	BookTitle     string `json:"bookTitle"`     // 书籍标题
	BookPic       string `json:"bookPic"`       // 书籍图片
	DownloadLink  string `json:"downloadLink"`  // 下载链接
	BookIntro     string `json:"bookIntro"`     // 书籍介绍
	ScanNumber    int    `json:"scanNumber"`    // 浏览数
	CommentNumber int    `json:"commentNumber"` // 评论数
	OnShelfDate   string `json:"onShelfDate"`   // 上架日期
	RequirePoints int    `json:"requirePoints"` // 下载所需积分
	TypeId        int    `json:"typeId"`        //  所属分类
}
