package cruise

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

type Service struct {
	store        *store.Store
	notifier     face.Notifier
	clock        func() time.Time
	defaultGrace time.Duration
}
