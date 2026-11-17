package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"

func (o *Tester) Service() *service.Service {
	return o.service
}
