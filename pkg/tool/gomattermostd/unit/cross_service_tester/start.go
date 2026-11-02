package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/worker"
)

func (o *Tester) Start() {
	for _, w := range o.Service.Workers() {
		w.Start()
		o.t.Cleanup(w.Close)
	}

	k := worker.NewEvery(
		o.Service.FlushDue,
		logger.New(o.t.Context()),
		o.Reporter,
		constant.FlushTick,
	)
	k.Start()
	o.t.Cleanup(k.Close)
}
