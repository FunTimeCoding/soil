package service

import "github.com/funtimecoding/soil/pkg/netbox/journal_entry"

func (s *Service) AddVirtualJournalEntry(
	machine string,
	kind string,
	comments string,
) (*journal_entry.Entry, error) {
	return s.client.AddVirtualJournalEntry(machine, kind, comments)
}
