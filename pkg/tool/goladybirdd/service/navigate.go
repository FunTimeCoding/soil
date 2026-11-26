package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) Navigate(
	tab string,
	locator string,
) error {
	if locator == "" {
		return validation.New("url is required")
	}

	t, e := s.resolveTab(tab)

	if e != nil {
		return e
	}

	return s.attempt(
		func() error {
			return s.client.Navigate(t, locator)
		},
	)
}
