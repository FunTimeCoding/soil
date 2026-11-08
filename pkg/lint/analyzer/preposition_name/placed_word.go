package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
)

func placedWord(name string) string {
	parts := segment.Segments(name)
	last := len(parts) - 1

	for i, s := range parts {
		if constant.PrepositionAnywhere[s] {
			return s
		}

		if constant.PrepositionNotFirst[s] &&
			(i > 0 || last == 0) &&
			parts[0] != constant.NewSegment {
			return s
		}

		if constant.PrepositionLast[s] && i == last {
			return s
		}
	}

	return ""
}
