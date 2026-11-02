package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/option"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
)

func New(o *option.Jellyfin) *Daemon {
	c := o.Source()

	return &Daemon{source: c, service: service.New(c)}
}
