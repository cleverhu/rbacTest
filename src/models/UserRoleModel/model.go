package UserRoleModel

type UserRoleResult struct {
	ID       int    `gorm:"column:uid" json:"uid"`
	UserName string `gorm:"column:username" json:"username"`
	RoleName string `gorm:"column:rolename" json:"role_name"`
}

func NewUserRoleResult() *UserRoleResult {
	return &UserRoleResult{}
}
