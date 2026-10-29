package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) Play(
	sessionIdentifier string,
	itemIdentifiers []string,
	playCommand string,
	startPositionTicks int64,
) error {
	if sessionIdentifier == "" {
		return validation.New("session_id is required")
	}

	if len(itemIdentifiers) == 0 {
		return validation.New("item_ids is required")
	}

	if playCommand == "" {
		playCommand = "PlayNow"
	}

	return s.source.Play(
		sessionIdentifier,
		itemIdentifiers,
		playCommand,
		startPositionTicks,
	)
}
