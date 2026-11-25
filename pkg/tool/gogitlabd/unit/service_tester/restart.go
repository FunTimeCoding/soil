package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"

func (o *Tester) Restart() *service.Service {
	o.t.Helper()
	o.service = nil

	return o.Service()
}
