package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/service"
)

func New(o *option.Firefox) *Daemon {
	c := o.Client()

	return &Daemon{client: c, service: service.New(c, o.DownloadDirectory)}
}
