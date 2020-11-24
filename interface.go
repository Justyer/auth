package auth

type IAuth interface {
	// 添加用户
	AddUser(IUser) (int64, error)
	// 添加角色
	AddRole(IRole) (int64, error)
	// 添加权限
	AddPerm(IPerm) (int64, error)
	// 添加菜单
	AddMenu(IMenu) (int64, error)
	// 修改用户角色
	UserRoleLink(IUserRole) error
	// // 修改角色权限
	RolePermLink(IRolePerm) error
}

type IUser interface {
	Add() (int64, error)
	Del() error
}

type IRole interface {
	Add() (int64, error)
	Del() error
}

type IPerm interface {
	Add() (int64, error)
	Del() error
}

type IMenu interface {
	Add() (int64, error)
	Del() error
}

type IUserRole interface {
	Add() error
	Del() error
}

type IRolePerm interface {
	Add() error
	Del() error
}
