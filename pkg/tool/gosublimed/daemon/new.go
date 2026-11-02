package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/option"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/service"
)

func New(o *option.Sublime) *Daemon {
	c := o.Source()

	return &Daemon{source: c, service: service.New(c)}
}
