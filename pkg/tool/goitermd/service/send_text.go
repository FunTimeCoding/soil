package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) SendText(
	identifier string,
	text string,
	enter bool,
) error {
	if text == "" {
		return validation.New("text is required")
	}

	return s.attempt(
		func() error {
			if e := s.client.SendText(identifier, text); e != nil {
				return e
			}

			if enter {
				return s.client.SendKey(identifier, "enter")
			}

			return nil
		},
	)
}
