package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) VirtualJournalEntries(
	machine string,
	limit int32,
	offset int32,
) ([]*journal_entry.Entry, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*journal_entry.Entry, error) {
			return s.client.VirtualJournalEntries(machine, limit, offset)
		},
	)
}
