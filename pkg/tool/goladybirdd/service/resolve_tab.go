package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) resolveTab(identifier string) (string, error) {
	tabs, e := s.Tabs()

	if e != nil {
		return "", e
	}

	if len(tabs) == 0 {
		return "", validation.New("browser has no tab open")
	}

	if identifier == "" {
		return tabs[0].Identifier, nil
	}

	for _, t := range tabs {
		if t.Identifier == identifier {
			return identifier, nil
		}
	}

	return "", not_found.New("tab", identifier)
}
