package storage

import (
	types "github.com/coreos/fcct/base/v0_2"
	"github.com/coreos/fcct/config/v1_1"
	"github.com/coreos/ignition/v2/config/util"
)

type FccbFile struct {
	file types.File
}

func NewFile(
	path string,
	overwrite bool,
	file []byte,
	append bool,
	mode int,
	owner *string,
	group *string) *FccbFile {

	f := types.File{
		Path:      path,
		Group:     types.NodeGroup{Name: group},
		Overwrite: &overwrite,
		User:      types.NodeUser{Name: owner},
		Mode:      &mode,
	}

	resource := types.Resource{
		Inline:       util.StrToPtr(string(file)),
		Verification: types.Verification{},
	}
	if append {
		f.Append = []types.Resource{
			resource,
		}
	} else {
		f.Contents = resource
	}

	return &FccbFile{
		file: f,
	}
}

func (f FccbFile) Add(fcc *v1_1.Config) {
	fcc.Storage.Files = append(fcc.Storage.Files, f.file)
}

func (f FccbFile) Remove(fcc *v1_1.Config) {
	panic("implement me")
}

func (f FccbFile) Replace(fcc *v1_1.Config) {
	panic("implement me")
}

