package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"time"
)

func New(
	s *store.Store,
	l *logger.Logger,
	predicate string,
	retention time.Duration,
	clock func() time.Time,
	command func(s ...string) (string, error),
	stream func(s ...string) face.Stream,
) *Service {
	return &Service{
		store:     s,
		logger:    l,
		predicate: predicate,
		retention: retention,
		clock:     clock,
		command:   command,
		stream:    stream,
		last:      make(map[string]string),
	}
}
