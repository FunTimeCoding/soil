package service

import "github.com/funtimecoding/soil/pkg/netbox/journal_entry"

func (s *Service) UpdateJournalEntry(
	identifier int32,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	return s.client.UpdateJournalEntry(identifier, kind, comments)
}
