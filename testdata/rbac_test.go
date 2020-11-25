package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/Justyer/auth"
	"github.com/Justyer/auth/rbac"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("root:495495@tcp(127.0.0.1:3306)/db_groot?charset=utf8&parseTime=True"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestUser(t *testing.T) {
	cfg := auth.Config{DB: db}

	mgr := auth.NewRBAC()

	u := &rbac.User{
		Config: cfg,
		Nick:   "川桑",
		Name:   "dengxc",
		Phone:  "17604626207",
		Pass:   "123",
	}
	uid, err := mgr.AddUser(u)
	fmt.Println(uid, err)
}
func TestRole(t *testing.T) {
	cfg := auth.Config{DB: db}

	mgr := auth.NewRBAC()

	u := &rbac.Role{
		Config: cfg,
		Name:   "KAMISAMA",
	}
	rid, err := mgr.AddRole(u)
	fmt.Println(rid, err)
}

func TestUserRole(t *testing.T) {
	cfg := auth.Config{DB: db}

	mgr := auth.NewRBAC()

	ur := &rbac.UserRole{
		Config: cfg,
		UserId: 1,
		RoleId: 1,
	}
	err := mgr.UserRoleLink(ur)
	fmt.Println(err)
}
