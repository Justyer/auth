package impl

import (
	"github.com/Justyer/auth"
	"gorm.io/gorm/clause"
)

type UserRole struct {
	auth.Config

	UserId int64  `gorm:"column:user_id,primary_key"`
	RoleId int64  `gorm:"column:role_id,primary_key"`
	Status string `gorm:"column:status;default"`
}

func (UserRole) TableName() string {
	return "auth_user_role"
}

func (self *UserRole) Add() error {
	return self.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "role_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(self).Error
}

func (self *UserRole) Del() error {
	if self.UserId&self.RoleId == 0 {
		return self.Err.Msg("user_id not empty")
	}
	return self.DB.Table(self.TableName()).Where("user_id=? and role_id=?", self.UserId, self.RoleId).UpdateColumn("status", "disable").Error
}
