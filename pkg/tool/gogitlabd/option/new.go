package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
)

func New() *Gitlab {
	return &Gitlab{
		Client:        BuildClient,
		Connector:     func() face.Client { return nil },
		WatchInterval: constant.WatchInterval,
		SweepInterval: constant.SweepInterval,
	}
}
