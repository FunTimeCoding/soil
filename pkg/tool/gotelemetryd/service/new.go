package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"time"
)

func New(
	s *store.Store,
	clock func() time.Time,
	r face.Reporter,
) *Service {
	return &Service{store: s, clock: clock, reporter: r}
}
