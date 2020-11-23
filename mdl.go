package auth

type User struct {
	UID   int64  `gorm:"column:uid,primary_key" json:"uid"`
	Unick string `gorm:"column:unick;default:''" json:"unick"`
	Uname string `gorm:"column:uname;default:''" json:"uname"`
	Upass string `gorm:"column:upass;default:''" json:"upass"`
}

type Role struct {
	RID   int64  `gorm:"column:rid,primary_key" json:"rid"`
	Rname string `gorm:"column:rname;default:''" json:"rname"`
}

type Perm struct {
	PID   int64  `gorm:"column:pid,primary_key" json:"pid"`
	Pname string `gorm:"column:pname;default:''" json:"pname"`
}

type Menu struct {
	MID   int64  `gorm:"column:mid,primary_key" json:"mid"`
	Mname string `gorm:"column:mname;default:''" json:"mname"`
}

type All struct {
	User
	Role
	Perm
	Menu
}
