package sweeper

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"sync"
	"time"
)

type Sweeper struct {
	store     *store.Store
	notifier  *notifier.Notifier
	forgetter face.Forgetter
	logger    *logger.Logger
	reporter  soilFace.Reporter
	window    time.Duration
	interval  time.Duration
	done      chan struct{}
	running   bool
	mutex     sync.Mutex
}
