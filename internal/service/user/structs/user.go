package structs

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Email    string `json:"email"`
	IsAdmin  bool
}
