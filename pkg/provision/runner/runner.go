package runner

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
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
	registry        face.ProcessRegistry
	logger          *logger.Logger
	recovery        *recovery.Recovery
	trigger         chan TriggerRequest
	sync            chan SyncRequest
	stop            chan struct{}
}
