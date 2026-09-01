package service

import (
	keepassConstant "github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"
	"github.com/tobischo/gokeepasslib/v3"
	"strings"
)

func (s *Service) Search(query string) []*credential.Credential {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	needle := strings.ToLower(query)
	var result []*credential.Credential
	s.client.Walk(
		func(
			path string,
			_ *gokeepasslib.Group,
			entry *gokeepasslib.Entry,
		) {
			haystack := strings.ToLower(
				strings.Join(
					[]string{
						entry.GetTitle(),
						entry.GetContent(keepassConstant.UserNameKey),
						entry.GetContent(constant.LocatorKey),
						entry.GetContent(constant.NotesKey),
						path,
					},
					"\n",
				),
			)

			if strings.Contains(haystack, needle) {
				result = append(result, describe(path, entry))
			}
		},
	)

	return result
}
