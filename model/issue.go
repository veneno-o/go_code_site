package model

type Issue struct {
	Id            int    `json:"id"`
	IssueTitle    string `json:"issueTitle"`    // 问题标题
	IssueContent  string `json:"issueContent"`  // 问题描述
	IssuePic      string `json:"issuePic"`      // 问题图片
	ScanNumber    int    `json:"scanNumber"`    //	问题浏览量
	CommentNumber int    `json:"commentNumber"` //	评论数
	IssueStatus   bool   `json:"issueStatus"`   //	问题状态
	IssueDate     string `json:"issueDate"`     //	问题时间
	UserId        int    `json:"userId"`        //	用户 id
	TypeId        int    `json:"typeId"`
}
