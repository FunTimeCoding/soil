//go:build local

package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"testing"
	"time"
)

func TestSetupReturnsFalseSkipsApply(t *testing.T) {
	applied := false
	r := runner.New(
		runner.Configuration{
			SetupFunction: func() bool { return false },
			ApplyFunction: func(
				_ map[string]any,
				_ string,
			) any {
				applied = true

				return nil
			},
		},
		logger.New(context.Background()),
		memory.New(),
	)
	r.Start()
	defer r.Stop()
	time.Sleep(100 * time.Millisecond)
	assert.False(t, applied)
}
