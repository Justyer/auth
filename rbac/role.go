package rbac

import "github.com/Justyer/auth"

type Role struct {
	Config `gorm:"-" json:"-"`

	RoleId int64  `gorm:"column:role_id;primaryKey" json:"role_id,omitempty"`
	Name   string `gorm:"column:name;default:''" json:"name,omitempty"`
	Status string `gorm:"column:status;default:'enable'" json:"status,omitempty"`
}

func (Role) TableName() string {
	return "auth_role"
}

func (self *Role) Add() (rid int64, err error) {
	switch {
	case self.RoleId != 0:
		err = self.DB.Where("role_id=?", self.RoleId).Updates(self).Error
	case self.Name != "":
		err = self.DB.Create(self).Error
	default:
		err = self.Err.Msg("role_id|role_name not empty")
	}
	rid = self.RoleId
	return
}

func (self *Role) Get() (auth.IRole, error) {
	err := self.DB.Where("role_id=?", self.RoleId).Find(self).Error
	return self, err
}

func (self *Role) Del() (err error) {
	switch {
	case self.RoleId != 0:
		err = self.DB.Where("role_id=?", self.RoleId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("role_id not empty")
	}
	return
}
