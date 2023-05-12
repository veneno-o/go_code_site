package model

type Interview struct {
	Id               int    `json:"id"`
	InterviewTitle   string `json:"interviewTitle"`   // 面试题标题
	InterviewContent string `json:"interviewContent"` // 面试题内容
	OnShelfDate      string `json:"onShelfDate"`      // 上架日期
	TypeId           string `json:"typeId"`           //  所属分类
}
