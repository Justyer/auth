package auth

type IAuth interface {
	// 添加用户
	AddUser(IUser) error
	// 添加角色
	AddRole(IRole) error
	// 添加权限
	AddPerm(IPerm) error
	// 添加菜单
	AddMenu(IMenu) error
	// 修改用户角色
	// UserRoleLink(int64, int64) error
	// // 修改角色权限
	// RolePermLink(int64, int64) error
}

type IUser interface {
	Add() error
	Delete() error
}

type IRole interface {
	Add() error
	Delete() error
}

type IPerm interface {
	Add() error
	Delete() error
}

type IMenu interface {
	Add() error
	Delete() error
}
