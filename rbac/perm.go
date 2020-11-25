package rbac

import "github.com/Justyer/auth"

type Perm struct {
	auth.Config `gorm:"-"`

	PermId int64  `gorm:"column:perm_id,primary_key"`
	Name   string `gorm:"column:name;default"`
	Status string `gorm:"column:status;default"`
}

func (Perm) TableName() string {
	return "auth_perm"
}

func (self *Perm) Add() (pid int64, err error) {
	pid = self.PermId
	switch {
	case self.PermId != 0:
		err = self.DB.Where("perm_id=?", self.PermId).Updates(self).Error
	case self.Name != "":
		err = self.DB.Create(self).Error
	default:
		err = self.Err.Msg("perm_id|perm_name not empty")
	}
	return
}

func (self *Perm) Del() (err error) {
	switch {
	case self.PermId != 0:
		err = self.DB.Where("perm_id=?", self.PermId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("perm_id not empty")
	}
	return
}
