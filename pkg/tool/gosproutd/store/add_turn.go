package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/turn"
)

func (s *Store) AddTurn(
	identifier uint,
	author constant.Author,
	content string,
) *turn.Turn {
	t := turn.New(identifier, author, content)
	errors.PanicOnError(s.mapper.Create(t).Error)

	return t
}
