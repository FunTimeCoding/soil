package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/option"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/service"
)

func New(o *option.Chrome) *Daemon {
	c := o.Source()

	return &Daemon{source: c, service: service.New(c, o.DownloadDirectory)}
}
