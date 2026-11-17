package service_tester

import "github.com/funtimecoding/soil/pkg/tool/goagentd/service"

func (o *Tester) Service() *service.Service {
	if o.service == nil {
		o.service = service.New(o.Runner)
	}

	return o.service
}
