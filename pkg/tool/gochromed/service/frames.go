package service

import (
	chromiumConstant "github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
)

func (s *Service) frames(identifier string) []*tab.Tab {
	tabs, e := guarded(s.guard, s.source.Tabs)

	if e != nil {
		return nil
	}

	var result []*tab.Tab

	for _, t := range tabs {
		if t.Type == chromiumConstant.IframeTabType &&
			t.ParentIdentifier == identifier {
			result = append(result, t)
		}
	}

	return result
}
