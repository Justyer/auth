package impl

import "github.com/Justyer/auth"

type Role struct {
	auth.Config

	RoleId int64  `gorm:"column:role_id,primary_key"`
	Name   string `gorm:"column:name;default"`
	Status string `gorm:"column:status;default"`
}

func (Role) TableName() string {
	return "auth_role"
}

func (self *Role) Add() (rid int64, err error) {
	if self.Name != "" {
		err = self.DB.Table(self.TableName()).Create(self).Error
		rid = self.RoleId
		return
	}
	return 0, self.Err.Msg("role_id|name not empty")
}

func (self *Role) Delete() error {
	if self.RoleId != 0 {
		return self.DB.Table(self.TableName()).Where("role_id=?", self.RoleId).UpdateColumn("status", "disable").Error
	}
	return self.Err.Msg("role_id not empty")
}
