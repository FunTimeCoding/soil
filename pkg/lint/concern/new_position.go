package concern

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/token"
)

func NewPosition(
	key string,
	text string,
	position token.Position,
) *Concern {
	return &Concern{
		Key:  key,
		Text: text,
		Path: position.Filename,
		Type: constant.ConcernLine,
		Line: position.Line,
	}
}
