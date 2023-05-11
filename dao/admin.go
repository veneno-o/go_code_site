package dao

type Admin struct {
	Id         int    `json:"id"`
	LoginId    string `json:"loginId"`    // 账号
	LoginPwd   string `json:"loginPwd"`   // 密码
	Nickname   string `json:"nickname"`   //昵称
	Avatar     string `json:"avatar"`     // 头像
	Permission int    `json:"permission"` // 权限
	Enabled    bool   `json:"enabled"`    // 是否可用
}
