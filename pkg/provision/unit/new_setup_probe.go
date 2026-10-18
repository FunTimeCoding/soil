package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"sync/atomic"
)

func newSetupProbe(setup *atomic.Bool) *runner.Runner {
	return runner.New(
		runner_option.Option{
			SetupFunction: func(context.Context) bool {
				setup.Store(true)

				return false
			},
		},
		logger.New(context.Background()),
		memory.New(),
	)
}
