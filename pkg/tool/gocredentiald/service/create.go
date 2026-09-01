package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/tobischo/gokeepasslib/v3"
)

func (s *Service) Create(
	groupPath string,
	title string,
	fields map[string]string,
) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	group := s.client.GroupByPath(groupPath)

	if group == nil {
		return "", not_found.New("group", groupPath)
	}

	entry := gokeepasslib.NewEntry()
	applyFields(&entry, map[string]string{constant.TitleKey: title})
	applyFields(&entry, fields)
	s.touch(&entry)
	group.Entries = append(group.Entries, entry)

	return fmt.Sprintf("%x", entry.UUID), s.persist()
}
