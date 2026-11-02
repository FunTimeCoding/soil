package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
)

func New(o *option.Iterm) *Daemon {
	c := o.Source()

	return &Daemon{source: c, service: service.New(c, o.KeyInterval)}
}
