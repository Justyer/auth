package rbac

type Query struct {
	Config

	UserId int
	Users  []*User
	Roles  []*Role
	Perms  []*Perm
}

func (self *Query) GetAllByUser() (err error) {

	return
}
