package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
)

type Runner struct {
	repository    string
	clonePath     string
	toolPath      string
	applyFunction func(
		parameters map[string]any,
		triggerSource string,
	) any
	initFunction    func()
	setupFunction   func() bool
	cleanupFunction func()
	changeFunction  func(value any) []string
	downstream      []face.Downstream
	registry        face.ProcessRegistry
	logger          *logger.Logger
	reporter        face.Reporter
	recovery        *recovery.Recovery
	syncFailures    int
	trigger         chan trigger.Request
	sync            chan update.Request
	stop            chan struct{}
}
