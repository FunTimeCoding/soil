package service

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"strings"
)

func (s *Service) resolveTab(v *selector.Selector) (*tab.Tab, error) {
	tabs, e := s.source.Tabs()

	if e != nil {
		return nil, e
	}

	if v.Identifier != "" {
		for _, t := range tabs {
			if t.Identifier == v.Identifier {
				return t, nil
			}
		}

		return nil, not_found.New("tab", v.Identifier)
	}

	if v.Title != "" {
		for _, t := range tabs {
			if strings.Contains(t.Title, v.Title) {
				return t, nil
			}
		}

		return nil, not_found.Format(
			"no tab with title containing: %s",
			v.Title,
		)
	}

	if v.Locator != "" {
		for _, t := range tabs {
			if strings.Contains(t.Locator, v.Locator) {
				return t, nil
			}
		}

		return nil, not_found.Format(
			"no tab with URL containing: %s",
			v.Locator,
		)
	}

	for _, t := range tabs {
		if t.Type == constant.PageTabType {
			return t, nil
		}
	}

	return nil, not_found.Format("no page tabs open")
}
