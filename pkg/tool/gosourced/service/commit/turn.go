package commit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/turn_grant"
)

func (s *Service) Turn(x context.Context) (*turn_grant.Grant, error) {
	return s.turns.Take(x, turn.Session(x))
}
