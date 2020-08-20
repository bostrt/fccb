package pkg

import "github.com/coreos/fcct/config/v1_1"

type FccbThingv1_1 interface {
	Add(fcc *v1_1.Config)
	Remove(fcc *v1_1.Config)
	Replace(fcc *v1_1.Config)
}
