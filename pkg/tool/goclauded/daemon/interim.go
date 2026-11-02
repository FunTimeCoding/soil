package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/worker"
	"time"
)

func interim(
	tick func(),
	l *logger.Logger,
	r face.Reporter,
	interval time.Duration,
) *worker.Worker {
	return worker.NewEvery(
		func() []error {
			tick()

			return nil
		},
		l,
		r,
		interval,
	)
}
