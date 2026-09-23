package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) Answer(
	identifier uint,
	kind constant.AnswerKind,
	channel constant.AnswerChannel,
	value string,
) {
	errors.PanicOnError(
		s.mapper.Model(decision.Stub()).Where(
			"identifier = ?",
			identifier,
		).Updates(
			map[string]any{
				"state":          constant.StateAnswered,
				"answer_kind":    kind,
				"answer_channel": channel,
				"answer":         value,
				"answered_at":    s.clock(),
			},
		).Error,
	)
}
