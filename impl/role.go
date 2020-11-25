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
	rid = self.RoleId
	switch {
	case self.RoleId != 0:
		err = self.DB.Table(self.TableName()).Where("role_id=?", self.RoleId).Updates(self).Error
	case self.Name != "":
		err = self.DB.Table(self.TableName()).Create(self).Error
	default:
		err = self.Err.Msg("role_id|role_name not empty")
	}
	return
}

func (self *Role) Del() (err error) {
	switch {
	case self.RoleId != 0:
		err = self.DB.Table(self.TableName()).Where("role_id=?", self.RoleId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("role_id not empty")
	}
	return
}
