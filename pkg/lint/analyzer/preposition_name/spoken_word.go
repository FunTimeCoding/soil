package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"go/types"
)

func spokenWord(o types.Object) string {
	parts := segment.Segments(o.Name())

	if len(parts) == 0 {
		return ""
	}

	if parts[0] == constant.QuestionSegment && question(o) {
		return placedWord(parts, 1)
	}

	return placedWord(parts, 0)
}
