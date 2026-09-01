package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"
	"github.com/tobischo/gokeepasslib/v3"
)

func (s *Service) List() []*credential.Credential {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	var result []*credential.Credential
	s.client.Walk(
		func(
			path string,
			_ *gokeepasslib.Group,
			entry *gokeepasslib.Entry,
		) {
			result = append(result, describe(path, entry))
		},
	)

	return result
}
