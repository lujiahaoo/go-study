package models

type User struct {
	UserName string `json:username`
	PassWord string `json:password`
}

type JwtToken struct {
	Token string `json:token`
}
