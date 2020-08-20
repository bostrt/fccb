package systemd

import (
	types "github.com/coreos/fcct/base/v0_2"
	"github.com/coreos/fcct/config/v1_1"
	ignutil "github.com/coreos/ignition/v2/config/util"
)

type FccbUnit struct {
	unit types.Unit
}

func NewUnit(unitName string, unitContents []byte, enabled bool, masked bool) *FccbUnit {
	unit := types.Unit{
		Contents: ignutil.StrToPtr(string(unitContents)),
		Name:     unitName,
		Mask: &masked,
		Enabled: &enabled,
	}

	return &FccbUnit{
		unit: unit,
	}
}

func (funit *FccbUnit) Add(fcc *v1_1.Config) {
	fcc.Systemd.Units = append(fcc.Systemd.Units, funit.unit)
}

func (funit *FccbUnit) Remove(fcc *v1_1.Config) {
	for _,u := range fcc.Systemd.Units {
		if u.Name == funit.unit.Name {
			// TODO
			return
		}
	}
}

func (funit *FccbUnit) Replace(fcc *v1_1.Config) {
	return
}

func (funit *FccbUnit) AddDropin(name string, contents []byte) {
	d := types.Dropin{
		Contents: ignutil.StrToPtr(string(contents)),
		Name: name,
	}
	funit.unit.Dropins = append(funit.unit.Dropins, d)
}

func checkUnitExists(needle types.Unit, haystack []types.Unit) bool {
	for _, u := range haystack {
		if u.Name == needle.Name {
			return true
		}
	}

	return false
}
