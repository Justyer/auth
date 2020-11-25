package test

import (
	"fmt"
	"log"
	"testing"
	"unsafe"

	"github.com/Justyer/auth"
	"github.com/Justyer/auth/rbac"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestRBAC(t *testing.T) {
	d, err := gorm.Open(mysql.Open("root:495495@tcp(127.0.0.1:3306)/db_groot?charset=utf8&parseTime=True"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal(err)
	}
	cfg := auth.Config{DB: d}

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

	// ur := &rbac.UserRole{
	// 	Config: cfg,
	// 	UserId: 1,
	// 	RoleId: 2,
	// 	Status: "disable",
	// }
	// err = mgr.UserRoleLink(ur)
	// fmt.Println(err)

	d.Migrator().CreateTable(&rbac.User{})
}

func TestT(t *testing.T) {
	var i int
	fmt.Printf("type: %T byte:%d\n", i, unsafe.Sizeof(i))
	var i32 int32
	fmt.Printf("type: %T byte:%d\n", i32, unsafe.Sizeof(i32))
	var i64 int64
	fmt.Printf("type: %T byte:%d\n", i64, unsafe.Sizeof(i64))
}
