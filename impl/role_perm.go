package impl

import (
	"github.com/Justyer/auth"
	"gorm.io/gorm/clause"
)

type RolePerm struct {
	auth.Config

	RoleId int64  `gorm:"column:role_id,primary_key"`
	PermId int64  `gorm:"column:perm_id,primary_key"`
	Status string `gorm:"column:status;default"`
}

func (RolePerm) TableName() string {
	return "auth_role_perm"
}

func (self *RolePerm) Add() error {
	return self.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "role_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(self).Error
}

func (self *RolePerm) Del() error {
	if self.RoleId&self.PermId == 0 {
		return self.Err.Msg("user_id not empty")
	}
	return self.DB.Table(self.TableName()).Where("role_id=? and perm_id=?", self.RoleId, self.PermId).UpdateColumn("status", "disable").Error
}
