package auth

type RBAC struct {
	Cfg Config
}

func NewRBAC() IAuth {
	return &RBAC{}
}

func (self *RBAC) AddUser(u IUser) error {
	return u.Add()
}

func (self *RBAC) AddRole(r IRole) error {
	return r.Add()
}

func (self *RBAC) AddPerm(p IPerm) error {
	return p.Add()
}

func (self *RBAC) AddMenu(m IMenu) error {
	return m.Add()
}
