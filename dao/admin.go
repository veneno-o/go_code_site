package dao

type admin struct {
	Id         int    `json:"id"`
	LoginId    string `json:"loginId"`
	LoginPwd   string `json:"loginPwd"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Permission int    `json:"permission"`
	Enabled    bool   `json:"enabled"`
}
