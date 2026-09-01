package service

import (
	"fmt"
	"github.com/tobischo/gokeepasslib/v3"
)

func (s *Service) pathOf(identifier string) string {
	result := ""
	s.client.Walk(
		func(
			path string,
			_ *gokeepasslib.Group,
			entry *gokeepasslib.Entry,
		) {
			if fmt.Sprintf("%x", entry.UUID) == identifier {
				result = path
			}
		},
	)

	return result
}
