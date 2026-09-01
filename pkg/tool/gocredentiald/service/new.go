package service

import (
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"time"
)

func New(
	path string,
	password string,
	revealedField []string,
	clock func() time.Time,
	l *logger.Logger,
) *Service {
	revealed := map[string]bool{}

	for _, key := range revealedField {
		revealed[key] = true
	}

	return &Service{
		path:     path,
		password: password,
		client:   keepass.New(path, password),
		revealed: revealed,
		clock:    clock,
		logger:   l,
	}
}
