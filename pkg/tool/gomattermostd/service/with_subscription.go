package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
)

func (s *Service) WithSubscription(
	n *notifier.Notifier,
	w *watcher.Watcher,
) *Service {
	s.notifier = n
	s.watcher = w

	return s
}
