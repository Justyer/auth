package auth

type IAuth interface {
	AddUser(IUser) (int64, error) // 添加用户
	AddRole(IRole) (int64, error) // 添加角色
	AddPerm(IPerm) (int64, error) // 添加权限

	UserRoleLink(IUserRole) error // 修改用户角色
	RolePermLink(IRolePerm) error // 修改角色权限

	DelUser(IUser) error // 禁用用户
	DelRole(IRole) error // 禁用角色
	DelPerm(IPerm) error // 禁用权限

	UserList(IPage) ([]IUser, error) // 用户列表
	RoleList(IPage) ([]IRole, error) // 角色列表
	PermList(IPage) ([]IPerm, error) // 权限列表

	GetAllByUser(IQuery) error // 根据用户id获取整个权限信息
}

type IUser interface {
	Add() (int64, error)
	Get() (IUser, error)
	Del() error
}

type IRole interface {
	Add() (int64, error)
	Get() (IRole, error)
	Del() error
}

type IPerm interface {
	Add() (int64, error)
	Get() (IPerm, error)
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

type IPage interface {
	UserPageTurn() ([]IUser, error)
	RolePageTurn() ([]IRole, error)
	PermPageTurn() ([]IPerm, error)
}

type IQuery interface {
	GetAllByUser() error
}
