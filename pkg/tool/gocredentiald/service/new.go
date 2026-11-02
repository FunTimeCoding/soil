package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/keepass"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"time"
)

func New(
	c *keepass.Client,
	path string,
	password string,
	revealedField []string,
	clock func() time.Time,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	revealed := map[string]bool{}

	for _, key := range revealedField {
		revealed[key] = true
	}

	return &Service{
		path:     path,
		password: password,
		client:   c,
		revealed: revealed,
		clock:    clock,
		logger:   l,
		reporter: r,
	}
}
