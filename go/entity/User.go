package entity

//数据库表明白定义
func (User) TableName() string {
	return "server_user"
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}
