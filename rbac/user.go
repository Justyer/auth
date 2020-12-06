package rbac

import "github.com/Justyer/auth"

type User struct {
	Config `gorm:"-" json:"-"`

	UserId int64  `gorm:"column:user_id;primaryKey" json:"user_id,omitempty"`
	Nick   string `gorm:"column:nick;default:''" json:"nick,omitempty"`
	Phone  string `gorm:"column:phone;default:''" json:"phone,omitempty"`
	Name   string `gorm:"column:name;default:''" json:"name,omitempty"`
	Pass   string `gorm:"column:pass;default:''" json:"pass,omitempty"`
	Status string `gorm:"column:status;default:'enable'" json:"status,omitempty"`
}

func (User) TableName() string {
	return "auth_user"
}

func (self *User) Add() (uid int64, err error) {
	switch {
	case self.UserId != 0:
		err = self.DB.Where("user_id=?", self.UserId).Updates(self).Error
	case self.Nick != "" && self.Name != "" && self.Pass != "":
		err = self.DB.Create(self).Error
	default:
		err = self.Err.Msg("user_nick|user_name|user_pass not empty")
	}
	uid = self.UserId
	return
}

func (self *User) Get() (auth.IUser, error) {
	err := self.DB.Where("user_id=?", self.UserId).Find(self).Error
	return self, err
}

func (self *User) Del() (err error) {
	switch {
	case self.UserId != 0:
		err = self.DB.Where("user_id=?", self.UserId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("user_id not empty")
	}
	return
}
