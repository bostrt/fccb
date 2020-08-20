package passwd

import (
	types "github.com/coreos/fcct/base/v0_2"
	"github.com/coreos/fcct/config/v1_1"
)

type FccbUser struct {
	user types.PasswdUser
}

func NewUser(
	name string,
	password_hash *string,
	uid *int,
	gecos *string,
	home_dir *string,
	no_create_home bool,
	primary_group *string,
	no_user_group bool,
	no_log_init bool,
	shell *string,
	system bool) *FccbUser {
	
	user := types.PasswdUser{
		Gecos:             gecos,
		HomeDir:           home_dir,
		Name:              name,
		NoCreateHome:      &no_create_home,
		NoLogInit:         &no_log_init,
		NoUserGroup:       &no_user_group,
		PasswordHash:      password_hash,
		PrimaryGroup:      primary_group,
		Shell:             shell,
		System:            &system,
		UID:               uid,
	}
	
	u := &FccbUser{
		user: user,
	}
	
	return u
}

func (f *FccbUser) Add(fcc *v1_1.Config) {
	fcc.Passwd.Users = append(fcc.Passwd.Users, f.user)
}

func (f *FccbUser) Remove(fcc *v1_1.Config) {
	panic("implement me")
}

func (f *FccbUser) Replace(fcc *v1_1.Config) {
	panic("implement me")
}

func (f *FccbUser) AddSSHAuthKey(key string) {
	authKey := types.SSHAuthorizedKey(key)
	f.user.SSHAuthorizedKeys = append(f.user.SSHAuthorizedKeys, authKey)
}

func (f *FccbUser) AddGroup(name string) {
	group := types.Group(name)
	f.user.Groups = append(f.user.Groups, group)
}