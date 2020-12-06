package rbac

import "github.com/Justyer/auth"

type Page struct {
	Config `json:"-"`

	Page     int   `json:"page,omitempty"`
	Count    int   `json:"count,omitempty"`
	Enable   bool  `json:"enable,omitempty"`
	Level    int32 `json:"level,omitempty"`     // 权限层级
	ParentId int64 `json:"parent_id,omitempty"` // 父级权限id
}

func (self *Page) UserPageTurn() (users []auth.IUser, err error) {
	var us []*User
	mid := self.DB.Limit(self.Count).Offset((self.Page - 1) * self.Count)
	if self.Enable {
		mid.Where("status='enable'")
	}
	err = mid.Find(&us).Error
	for i := 0; i < len(us); i++ {
		var uu auth.IUser = us[i]
		users = append(users, uu)
	}
	return
}

func (self *Page) RolePageTurn() (roles []auth.IRole, err error) {
	var rs []*Role
	mid := self.DB.Limit(self.Count).Offset((self.Page - 1) * self.Count)
	if self.Enable {
		mid.Where("status='enable'")
	}
	err = mid.Find(&rs).Error
	for i := 0; i < len(rs); i++ {
		var rr auth.IRole = rs[i]
		roles = append(roles, rr)
	}
	return
}

func (self *Page) PermPageTurn() (perms []auth.IPerm, err error) {
	var ps []*Perm
	mid := self.DB.Limit(self.Count).Offset((self.Page - 1) * self.Count)
	if self.Enable {
		mid.Where("status='enable'")
	}
	mid.Where("level=?", self.Level)
	if self.ParentId != 0 {
		mid.Where("parent_id=?", self.ParentId)
	}
	err = mid.Find(&ps).Error
	for i := 0; i < len(ps); i++ {
		var pp auth.IPerm = ps[i]
		perms = append(perms, pp)
	}
	return
}
