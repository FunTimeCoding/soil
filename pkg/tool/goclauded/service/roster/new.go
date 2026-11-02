package roster

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	messages *message.Service,
	sessions *session.Service,
	cache *session_cache.Cache,
	n face.Notifier,
	clock func() time.Time,
	l *logger.Logger,
) *Service {
	return &Service{
		store:    s,
		message:  messages,
		session:  sessions,
		cache:    cache,
		notifier: n,
		clock:    clock,
		logger:   l,
	}
}
