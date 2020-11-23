package impl

import "github.com/Justyer/auth"

type Menu struct {
	auth.Config

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
	if self.Name != "" && self.Path != "" && self.Level != 0 {
		err = self.DB.Table(self.TableName()).Create(self).Error
		mid = self.MenuId
		return
	}
	return 0, self.Err.Msg("menu_name|menu_path|menu_level not empty")
}

func (self *Menu) Delete() error {
	if self.MenuId != 0 {
		return self.DB.Table(self.TableName()).Where("menu_id=?", self.MenuId).UpdateColumn("status", "disable").Error
	}
	return self.Err.Msg("menu_id not empty")
}
