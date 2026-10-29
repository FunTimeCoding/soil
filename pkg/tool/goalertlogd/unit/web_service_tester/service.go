package web_service_tester

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"

func (o *Tester) Service() *service.Service {
	return o.server.Service
}
