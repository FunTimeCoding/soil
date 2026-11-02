package service

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"time"
)

func New(
	s *store.Store,
	m chat.MattermostClient,
	g *peer.Guard,
	l *logger.Logger,
	r face.Reporter,
	debounce time.Duration,
	purge time.Duration,
) *Service {
	return &Service{
		store:      s,
		mattermost: m,
		guard:      g,
		logger:     l,
		reporter:   r,
		debounce:   debounce,
		purge:      purge,
	}
}
