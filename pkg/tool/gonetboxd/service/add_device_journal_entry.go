package service

import "github.com/funtimecoding/soil/pkg/netbox/journal_entry"

func (s *Service) AddDeviceJournalEntry(
	device string,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	return guarded(
		s.guard,
		func() (*journal_entry.Entry, error) {
			return s.client.AddDeviceJournalEntry(device, kind, comments)
		},
	)
}
