package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/turn"
	"unicode/utf8"
)

func (s *Service) AddTurn(
	identifier uint,
	author constant.Author,
	content string,
) (*turn.Turn, error) {
	if length := utf8.RuneCountInString(
		content,
	); length > constant.MaximumTurnLength {
		return nil, fmt.Errorf(
			"turn is %d characters, the limit is %d",
			length,
			constant.MaximumTurnLength,
		)
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if author == constant.AuthorSession && s.store.TurnCount(
		identifier,
		author,
	) >= constant.MaximumTurnCount {
		return nil, fmt.Errorf(
			"this decision has used all %d session turns - clear it or open a new one under a new framing",
			constant.MaximumTurnCount,
		)
	}

	t := s.store.AddTurn(identifier, author, content)
	s.notifier.Notify()

	return t, nil
}
