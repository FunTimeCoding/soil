package service

import "github.com/funtimecoding/soil/pkg/netbox/journal_entry"

func (s *Service) DeviceJournalEntries(
	device string,
	limit int32,
	offset int32,
) ([]*journal_entry.Entry, error) {
	return s.client.DeviceJournalEntries(device, limit, offset)
}
