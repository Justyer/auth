package auth

type RBAC struct{}

func NewRBAC() IAuth {
	return &RBAC{}
}

func (self *RBAC) AddUser(u IUser) (int64, error) {
	return u.Add()
}

func (self *RBAC) AddRole(r IRole) (int64, error) {
	return r.Add()
}

func (self *RBAC) AddPerm(p IPerm) (int64, error) {
	return p.Add()
}

func (self *RBAC) UserRoleLink(ur IUserRole) error {
	return ur.Add()
}

func (self *RBAC) RolePermLink(rp IRolePerm) error {
	return rp.Add()
}

func (self *RBAC) DelUser(u IUser) error {
	return u.Del()
}

func (self *RBAC) DelRole(r IRole) error {
	return r.Del()
}

func (self *RBAC) DelPerm(p IPerm) error {
	return p.Del()
}

func (self *RBAC) UserList(p IPage) ([]IUser, error) {
	return p.UserPageTurn()
}

func (self *RBAC) RoleList(p IPage) ([]IRole, error) {
	return p.RolePageTurn()
}

func (self *RBAC) PermList(p IPage) ([]IPerm, error) {
	return p.PermPageTurn()
}

func (self *RBAC) GetAllByUser(q IQuery) error {
	return q.GetAllByUser()
}
