package embedder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service == nil {
		o.service = service_tester.New(
			o.t,
			service_tester.WithEmbedder(o.Environment()),
			service_tester.WithEmbedCounter(o.Counter()),
		).Service()
	}

	return o.service
}
