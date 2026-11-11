package cruise

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	n face.EventNotifier,
	clock func() time.Time,
) *Service {
	return &Service{store: s, notifier: n, clock: clock}
}
