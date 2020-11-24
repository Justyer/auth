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

func (self *RBAC) AddMenu(m IMenu) (int64, error) {
	return m.Add()
}

func (self *RBAC) UserRoleLink(ur IUserRole) error {
	return ur.Add()
}

func (self *RBAC) RolePermLink(rp IRolePerm) error {
	return rp.Add()
}
