package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	saltFace "github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"sync"
	"time"
)

type Runner struct {
	provision       *runner.Runner
	store           *store.Store
	salt            saltFace.SaltClient
	guard           *peer.Guard
	keySyncInterval time.Duration
	keySyncing      sync.WaitGroup
	logger          *logger.Logger
	reporter        face.Reporter
	recovery        *recovery.Recovery
}
