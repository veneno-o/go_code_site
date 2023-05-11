package dao

type User struct {
	Id            int    `json:"id"`            // mongodb 自动生成的 id
	LoginId       string `json:"loginId"`       // 账号
	LoginPwd      string `json:"loginPwd"`      // 密码
	Avatar        string `json:"avatar"`        // 头像
	Nickname      string `json:"nickname"`      // 昵称
	Mail          string `json:"mail"`          // 邮箱
	Qq            string `json:"qq"`            // QQ
	Wechat        string `json:"wechat"`        // 微信号
	Intro         string `json:"intro"`         // 个人介绍
	RegisterDate  string `json:"registerDate"`  // 注册时间
	LastLoginDate string `json:"lastLoginDate"` // 上次登录事件
	Points        int    `json:"points"`        // 积分
	Enabled       bool   `json:"enabled"`       // 是否可用
}
