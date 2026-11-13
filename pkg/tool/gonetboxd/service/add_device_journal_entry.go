package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/journal_entry"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) AddDeviceJournalEntry(
	device string,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*journal_entry.Entry, error) {
			return s.client.AddDeviceJournalEntry(device, kind, comments)
		},
	)
}
