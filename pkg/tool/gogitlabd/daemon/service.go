package daemon

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"

func (d *Daemon) Service() *service.Service {
	return d.service
}
