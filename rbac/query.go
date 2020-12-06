package rbac

type Query struct {
	Config

	UserId int64
	User   User
	Roles  []*Role
	Perms  []*Perm
}

func (self *Query) GetAllByUser() (err error) {
	err = self.DB.Where("user_id=?", self.UserId).Find(&self.User).Error
	if err != nil {
		return
	}
	sql := `select r.* from auth_user_role ur join auth_role r on ur.role_id=r.role_id where ur.user_id=? and ur.status='enable' and r.status='enable'`
	self.DB.Raw(sql, self.UserId).Find(&self.Roles)
	if len(self.Roles) == 0 {
		return
	}

	var rids []int64
	for _, r := range self.Roles {
		rids = append(rids, r.RoleId)
	}
	sql = `select p.* from auth_role_perm rp join auth_perm p on rp.perm_id=p.perm_id where rp.role_id in (?) and rp.status='enable' and p.status='enable'`
	self.DB.Raw(sql, rids).Find(&self.Perms)

	return
}
