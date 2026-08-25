package model_context

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"strings"
)

func (s *Server) resolveTab(
	tabIdentifier string,
	title string,
	l string,
) (*tab.Tab, error) {
	tabs := s.client.Tabs()

	if tabIdentifier != "" {
		for _, t := range tabs {
			if t.Identifier == tabIdentifier {
				return t, nil
			}
		}

		return nil, not_found.New("tab", tabIdentifier)
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
