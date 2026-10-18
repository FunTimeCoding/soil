package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
	"sync"
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
	setupFunction   func(context.Context) bool
	cleanupFunction func()
	changeFunction  func(value any) *change.Change
	downstream      []face.Downstream
	registry        library.ProcessRegistry
	logger          *logger.Logger
	reporter        library.Reporter
	recovery        *recovery.Recovery
	syncFailures    int
	trigger         chan trigger.Request
	sync            chan update.Request
	stop            chan struct{}
	context         context.Context
	cancel          context.CancelFunc
	mutex           sync.Mutex
	started         bool
	closed          bool
	finished        chan struct{}
}
