package label

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
)

func New(
	s *store.Store,
	n face.EventNotifier,
) *Service {
	return &Service{store: s, notifier: n}
}
