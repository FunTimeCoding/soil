package service

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
)

func (s *Service) frames(identifier string) []*tab.Tab {
	tabs, e := s.source.Tabs()

	if e != nil {
		return nil
	}

	var result []*tab.Tab

	for _, t := range tabs {
		if t.Type == constant.IframeTabType &&
			t.ParentIdentifier == identifier {
			result = append(result, t)
		}
	}

	return result
}
