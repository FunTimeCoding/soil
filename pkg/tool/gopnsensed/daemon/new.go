package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
)

func New(o *option.Opnsense) *Daemon {
	c := o.Source()

	return &Daemon{source: c, service: service.New(c)}
}
