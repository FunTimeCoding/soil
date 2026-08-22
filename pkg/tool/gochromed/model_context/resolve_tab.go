package model_context

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"strings"
)

func (s *Server) resolveTab(
	tabID string,
	title string,
	l string,
) (*tab.Tab, error) {
	tabs := s.client.Tabs()

	if tabID != "" {
		for _, t := range tabs {
			if t.Identifier == tabID {
				return t, nil
			}
		}

		return nil, not_found.New("tab", tabID)
	}

	if title != "" {
		for _, t := range tabs {
			if strings.Contains(t.Title, title) {
				return t, nil
			}
		}

		return nil, not_found.Format("no tab with title containing: %s", title)
	}

	if l != "" {
		for _, t := range tabs {
			if strings.Contains(t.Locator, l) {
				return t, nil
			}
		}

		return nil, not_found.Format("no tab with URL containing: %s", l)
	}

	for _, t := range tabs {
		if t.Type == constant.PageTabType {
			return t, nil
		}
	}

	return nil, not_found.Format("no page tabs open")
}
