package dao

type Comment struct {
	Id             int    `json:"id"`
	UserId         int    `json:"userId"`         //  所属用户
	IssueId        int    `json:"issueId"`        //  所属文章
	BookId         int    `json:"bookId"`         //  所属书籍
	TypeId         int    `json:"typeId"`         //  所属类型
	CommentContent string `json:"commentContent"` //  评论内容
	CommentDate    string `json:"commentDate"`    //  评论日期
	CommentType    int    `json:"commentType"`    //评论类型
}
