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

var (
	db  *gorm.DB
	cfg rbac.Config
	mgr auth.IAuth
)

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("root:495495@tcp(127.0.0.1:3306)/db_groot?charset=utf8&parseTime=True"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	cfg = rbac.Config{DB: db}

	mgr = auth.NewRBAC()
}

func TestUser(t *testing.T) {
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
	u := &rbac.Role{
		Config: cfg,
		Name:   "KAMISAMA",
	}
	rid, err := mgr.AddRole(u)
	fmt.Println(rid, err)
}

func TestUserRole(t *testing.T) {
	ur := &rbac.UserRole{
		Config: cfg,
		UserId: 1,
		RoleId: 1,
	}
	err := mgr.UserRoleLink(ur)
	fmt.Println(err)
}

func TestUserList(t *testing.T) {
	page := &rbac.Page{
		Config: cfg,
		Page:   1,
		Count:  5,
		Enable: true,
	}

	users, err := mgr.UserList(page)
	fmt.Println(users)
	fmt.Println(err)
}

func TestRoleList(t *testing.T) {
	page := &rbac.Page{
		Config: cfg,
		Page:   1,
		Count:  5,
		Enable: true,
	}

	roles, err := mgr.RoleList(page)
	fmt.Println(roles)
	fmt.Println(err)
}

func TestGetAllByUser(t *testing.T) {
	query := &rbac.Query{
		Config: cfg,
		UserId: 1,
	}

	err := mgr.GetAllByUser(query)
	fmt.Println("users", query.User)
	fmt.Println("roles", query.Roles)
	fmt.Println("perms", query.Perms)
	fmt.Println(err)
}
