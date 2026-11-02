package label

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
)

func New(
	s *store.Store,
	n face.Notifier,
) *Service {
	return &Service{store: s, notifier: n}
}
