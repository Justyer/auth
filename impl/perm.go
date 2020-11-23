package impl

import "github.com/Justyer/auth"

type Perm struct {
	auth.Config

	PermId int64  `gorm:"column:perm_id,primary_key"`
	Name   string `gorm:"column:name;default"`
	Status string `gorm:"column:status;default"`
}

func (Perm) TableName() string {
	return "auth_perm"
}

func (self *Perm) Add() (rid int64, err error) {
	if self.Name != "" {
		err = self.DB.Table(self.TableName()).Create(self).Error
		rid = self.PermId
		return
	}
	return 0, self.Err.Msg("perm_id|perm_name not empty")
}

func (self *Perm) Delete() error {
	if self.PermId != 0 {
		return self.DB.Table(self.TableName()).Where("perm_id=?", self.PermId).UpdateColumn("status", "disable").Error
	}
	return self.Err.Msg("perm_id not empty")
}
