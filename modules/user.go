package modules

type User struct {
	Id       int
	Username string
	Email    string
	AddTime  int
}

// 返回表名
func (User) TableName() string {
	return "gin_user"
}
