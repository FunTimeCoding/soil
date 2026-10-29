package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) SetVolume(
	sessionIdentifier string,
	level int,
) error {
	if sessionIdentifier == "" {
		return validation.New("session_id is required")
	}

	if level < 0 || level > 100 {
		return validation.New("level must be between 0 and 100")
	}

	return s.source.SetVolume(sessionIdentifier, level)
}
