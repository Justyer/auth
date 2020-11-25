package rbac

import (
	"github.com/Justyer/auth"
)

type Role struct {
	auth.Config `gorm:"-"`

	RoleId int64  `gorm:"column:role_id;type:int(11);autoIncrement;primaryKey"`
	Name   string `gorm:"column:name;type:varchar(10);not null;default:'';comment:角色"`
	Status string `gorm:"column:status;type:enum('enable','disable');not null;default:'enable';comment:状态"`

	// CreateTime int `gorm:"column:create_time;type:timestamp;not null;default:current_timestamp"`
	// UpdateTime int `gorm:"column:update_time;type:timestamp;autoUpdateTime"`
}

func (Role) TableName() string {
	return "auth_role"
}

func (self *Role) Add() (rid int64, err error) {
	rid = self.RoleId
	switch {
	case self.RoleId != 0:
		err = self.DB.Where("role_id=?", self.RoleId).Updates(self).Error
	case self.Name != "":
		err = self.DB.Create(self).Error
	default:
		err = self.Err.Msg("role_id|role_name not empty")
	}
	return
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
