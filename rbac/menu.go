package rbac

import "github.com/Justyer/auth"

type Menu struct {
	auth.Config `gorm:"-"`

	MenuId int64  `gorm:"column:menu_id,primary_key"`
	Name   string `gorm:"column:name;default"`
	Path   string `gorm:"column:path;default"`
	Level  int32  `gorm:"column:level;default"`
	Index  int32  `gorm:"column:index;default"`
	Status string `gorm:"column:status;default"`
}

func (Menu) TableName() string {
	return "auth_menu"
}

func (self *Menu) Add() (mid int64, err error) {
	switch {
	case self.MenuId != 0:
		err = self.DB.Where("menu_id=?", self.MenuId).Updates(self).Error
	case self.Name != "" && self.Path != "" && self.Level != 0:
		err = self.DB.Create(self).Error
	default:
		err = self.Err.Msg("menu_name|menu_path|menu_level not empty")
	}
	mid = self.MenuId
	return
}

func (self *Menu) Del() (err error) {
	switch {
	case self.MenuId != 0:
		err = self.DB.Where("menu_id=?", self.MenuId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("menu_id not empty")
	}
	return
}
