package service

import (
	"fmt"
	keepassConstant "github.com/funtimecoding/soil/pkg/keepass/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/entry_detail"
	"time"
)

func (s *Service) Get(identifier string) *entry_detail.Detail {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.refresh()
	entry, _ := s.client.EntryByIdentifier(identifier)

	if entry == nil {
		return nil
	}

	path := s.pathOf(identifier)
	fields := map[string]string{}

	for _, value := range entry.Values {
		switch value.Key {
		case constant.TitleKey:
			continue
		case keepassConstant.UserNameKey, constant.LocatorKey:
			fields[value.Key] = value.Value.Content

			continue
		}

		if s.revealed[value.Key] && !value.Value.Protected.Bool {
			fields[value.Key] = value.Value.Content

			continue
		}

		fields[value.Key] = constant.MaskedValue
	}

	modified := time.Time{}

	if entry.Times.LastModificationTime != nil {
		modified = entry.Times.LastModificationTime.Time
	}

	return entry_detail.New(
		fmt.Sprintf("%x", entry.UUID),
		path,
		entry.GetTitle(),
		fields,
		modified,
	)
}
