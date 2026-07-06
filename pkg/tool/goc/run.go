package goc

import (
	"github.com/funtimecoding/soil/pkg/ceph/goc"
	"github.com/funtimecoding/soil/pkg/tool/goc/option"
)

func Run(o *option.Ceph) {
	goc.Run(o.Command, false)
}
