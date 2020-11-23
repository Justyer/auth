package auth

type RBAC struct {
	Cfg Config
}

func NewRBAC() *RBAC {
	return &RBAC{}
}
