package passwd

import (
	types "github.com/coreos/fcct/base/v0_2"
	"github.com/coreos/fcct/config/v1_1"
)

type FccbGroup struct {
	group types.PasswdGroup
}

func NewGroup(name string, gid *int, password *string, system *bool) *FccbGroup {
	group := types.PasswdGroup{
		Gid:          gid,
		Name:         name,
		PasswordHash: password,
		System:       system,
	}

	g := &FccbGroup{
		group: group,
	}

	return g
}

func (f FccbGroup) Add(fcc *v1_1.Config) {
	fcc.Passwd.Groups = append(fcc.Passwd.Groups, f.group)
}

func (f FccbGroup) Remove(fcc *v1_1.Config) {
	panic("implement me")
}

func (f FccbGroup) Replace(fcc *v1_1.Config) {
	panic("implement me")
}
