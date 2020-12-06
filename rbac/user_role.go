package rbac

import (
	"gorm.io/gorm/clause"
)

type UserRole struct {
	Config `gorm:"-" json:"-"`

	UserId int64  `gorm:"column:user_id;primaryKey" json:"user_id,omitempty"`
	RoleId int64  `gorm:"column:role_id;primaryKey" json:"role_id,omitempty"`
	Status string `gorm:"column:status;default:'enable'" json:"status,omitempty"`
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

func (self *UserRole) Del() (err error) {
	switch {
	case self.UserId&self.RoleId != 0:
		err = self.DB.Where("user_id=? and role_id=?", self.UserId, self.RoleId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("user_id not empty")
	}
	return
}
