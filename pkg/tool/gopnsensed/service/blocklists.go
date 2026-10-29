package service

import "github.com/funtimecoding/soil/pkg/opnsense/blocklist"

func (s *Service) Blocklists(phrase string) ([]*blocklist.Blocklist, error) {
	return s.source.Blocklists(phrase)
}
