package service_tester

import "github.com/funtimecoding/soil/pkg/generative/model_context_server"

func (o *Tester) Served() *model_context_server.Server {
	o.t.Helper()

	if o.served == nil {
		o.served = model_context_server.Mount(o.t, o.Daemon(), o.Reporter)
	}

	return o.served
}
