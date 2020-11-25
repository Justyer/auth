package rbac

import (
	"github.com/Justyer/auth"
	"gorm.io/gorm/clause"
)

type RolePerm struct {
	auth.Config `gorm:"-"`

	RoleId int64  `gorm:"column:role_id,primary_key"`
	PermId int64  `gorm:"column:perm_id,primary_key"`
	Status string `gorm:"column:status;default"`
}

func (RolePerm) TableName() string {
	return "auth_role_perm"
}

func (self *RolePerm) Add() error {
	return self.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "role_id"}, {Name: "perm_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(self).Error
}

func (self *RolePerm) Del() (err error) {
	switch {
	case self.RoleId&self.PermId != 0:
		err = self.DB.Where("role_id=? and perm_id=?", self.RoleId, self.PermId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("user_id not empty")
	}
	return
}
