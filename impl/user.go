package impl

import "github.com/Justyer/auth"

type User struct {
	auth.Config

	UserId int64  `gorm:"column:user_id,primary_key"`
	Nick   string `gorm:"column:nick;default"`
	Phone  string `gorm:"column:phone;default"`
	Name   string `gorm:"column:name;default"`
	Pass   string `gorm:"column:pass;default"`
	Status string `gorm:"column:status;default"`
}

func (User) TableName() string {
	return "auth_user"
}

func (self *User) Add() (uid int64, err error) {
	uid = self.UserId
	switch {
	case self.UserId != 0:
		err = self.DB.Table(self.TableName()).Where("user_id=?", self.UserId).Updates(self).Error
	case self.Nick != "" && self.Name != "" && self.Pass != "":
		err = self.DB.Table(self.TableName()).Create(self).Error
	default:
		err = self.Err.Msg("user_nick|user_name|user_pass not empty")
	}
	return
}

func (self *User) Del() (err error) {
	switch {
	case self.UserId != 0:
		err = self.DB.Table(self.TableName()).Where("user_id=?", self.UserId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("user_id not empty")
	}
	return
}
