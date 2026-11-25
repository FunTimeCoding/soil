package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"github.com/funtimecoding/soil/pkg/watch/notifier"
)

func (s *Service) WithSubscription(
	n *notifier.Notifier,
	w *watcher.Watcher,
) *Service {
	s.notifier = n
	s.watcher = w

	return s
}
