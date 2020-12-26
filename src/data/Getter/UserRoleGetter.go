package Getter

import (
	"rbacTest/src/dbs"
	"rbacTest/src/models/UserRoleModel"
)

var UserRoleGetter IUserRole

func init() {
	UserRoleGetter = NewUserRoleImpl()
}

type IUserRole interface {
	UserRoleList() []*UserRoleModel.UserRoleResult
}

type UserRoleImpl struct {
}

func NewUserRoleImpl() *UserRoleImpl {
	return &UserRoleImpl{}
}

func (this *UserRoleImpl) UserRoleList() (urs []*UserRoleModel.UserRoleResult) {

	dbs.Orm.Raw("SELECT u.id as 'uid',u.name as 'username',r.name as 'rolename' from user as u,role as r, user_role as ur where u.id = ur.uid and r.id = ur.role_id ;").Find(&urs)

	return urs
}
