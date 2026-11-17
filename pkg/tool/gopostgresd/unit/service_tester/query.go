package service_tester

import "github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"

func (o *Tester) Query() *query.Service {
	o.t.Helper()

	if o.query == nil {
		o.query = query.New(o.peers())
	}

	return o.query
}
