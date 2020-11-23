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
	if self.Nick != "" && self.Name != "" && self.Pass != "" {
		err = self.DB.Table(self.TableName()).Create(self).Error
		uid = self.UserId
		return
	}
	return 0, self.Err.Msg("user_nick|user_name|user_pass not empty")
}

func (self *User) Delete() error {
	if self.UserId != 0 {
		return self.DB.Table(self.TableName()).Where("user_id=?", self.UserId).UpdateColumn("status", "disable").Error
	}
	return self.Err.Msg("user_id not empty")
}
