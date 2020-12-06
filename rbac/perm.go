package rbac

import "github.com/Justyer/auth"

type Perm struct {
	Config `gorm:"-" json:"-"`

	PermId   int64  `gorm:"column:perm_id;primaryKey" json:"perm_id,omitempty"`
	Name     string `gorm:"column:name;default:''" json:"name,omitempty"`
	Path     string `gorm:"column:path;default:''" json:"path,omitempty"`
	Level    int32  `gorm:"column:level;default:0" json:"level,omitempty"`
	ParentId int64  `gorm:"column:parent_id;default:0" json:"parent_id,omitempty"`
	Status   string `gorm:"column:status;default:'enable'" json:"status,omitempty"`
}

func (Perm) TableName() string {
	return "auth_perm"
}

func (self *Perm) Add() (pid int64, err error) {
	switch {
	case self.PermId != 0:
		self.DB.Where("perm_id=?", self.PermId).Updates(self)
	case self.Name != "" && self.Path != "" && self.Level != 0:
		err = self.DB.Create(self).Error
	default:
		err = self.Err.Msg("perm->[name|path|level] not empty")
	}
	pid = self.PermId
	return
}

func (self *Perm) Get() (auth.IPerm, error) {
	err := self.DB.Where("perm_id=?", self.PermId).Find(self).Error
	return self, err
}

func (self *Perm) Del() (err error) {
	switch {
	case self.PermId != 0:
		err = self.DB.Where("perm_id=?", self.PermId).UpdateColumn("status", "disable").Error
	default:
		err = self.Err.Msg("perm->[id] not empty")
	}
	return
}
