package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Tabs() ([]*tab.Tab, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*tab.Tab, error) {
			return s.source.Tabs()
		},
	)
}
