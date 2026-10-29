package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) SetTabTitle(
	tabIdentifier string,
	title string,
) error {
	if title == "" {
		return validation.New("title is required")
	}

	return s.source.SetTabTitle(tabIdentifier, title)
}
