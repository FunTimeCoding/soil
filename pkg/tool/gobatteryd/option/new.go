package option

import (
	"github.com/funtimecoding/soil/pkg/nut"
	"github.com/funtimecoding/soil/pkg/nut/face"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/inventory/instance"
)

func New() *Exporter {
	return &Exporter{
		Source: func(i *instance.Instance) face.Client {
			return nut.New(i.Host, i.ServerPort())
		},
	}
}
