package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/opnsense/blocklist"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Blocklists(phrase string) ([]*blocklist.Blocklist, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*blocklist.Blocklist, error) {
			return s.client.Blocklists(phrase)
		},
	)
}
