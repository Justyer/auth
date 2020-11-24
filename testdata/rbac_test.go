package test

import (
	"log"
	"testing"

	"github.com/Justyer/auth"
	"github.com/Justyer/auth/impl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRBAC(t *testing.T) {
	d, err := gorm.Open(mysql.Open("root:495495@tcp(127.0.0.1:3306)/db_groot?parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	d.LogMode(true)
	cfg := auth.Config{
		DB: d,
	}
	u := &impl.User{
		Config: cfg,
		Nick:   "川桑",
		Name:   "dengxc",
		Phone:  "17604626207",
		Pass:   "123",
	}
	rbac := auth.NewRBAC()
	rbac.AddUser(u)
}
