package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
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

		return nil, fmt.Errorf("tab not found: %s", tabID)
	}

	if title != "" {
		for _, t := range tabs {
			if strings.Contains(t.Title, title) {
				return t, nil
			}
		}

		return nil, fmt.Errorf("no tab with title containing: %s", title)
	}

	if l != "" {
		for _, t := range tabs {
			if strings.Contains(t.Locator, l) {
				return t, nil
			}
		}

		return nil, fmt.Errorf("no tab with URL containing: %s", l)
	}

	for _, t := range tabs {
		if t.Type == constant.PageTabType {
			return t, nil
		}
	}

	return nil, fmt.Errorf("no page tabs open")
}
