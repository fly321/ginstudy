package modules

type User struct {
	Id       int
	Username int
	Email    int
	AddTime  int
}

// 返回表名
func (User) TableName() string {
	return "gin_user"
}
