package service

import (
	keepassConstant "github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/tobischo/gokeepasslib/v3"
)

func (s *Service) LoadGroup(name string) map[string]string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	group := s.client.GroupByName(constant.EnvironmentGroup)

	if group == nil {
		return nil
	}

	var found *gokeepasslib.Entry

	for i := range group.Entries {
		if group.Entries[i].GetTitle() == name {
			found = &group.Entries[i]
		}
	}

	if found == nil {
		return nil
	}

	result := map[string]string{}

	for _, value := range found.Values {
		switch value.Key {
		case constant.TitleKey,
			keepassConstant.UserNameKey,
			constant.PasswordKey,
			constant.LocatorKey,
			constant.NotesKey:

			continue
		}

		result[value.Key] = value.Value.Content
	}

	return result
}
