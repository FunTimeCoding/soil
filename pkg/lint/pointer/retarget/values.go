package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
	"unicode/utf8"
)

func Values(
	moves []*path_move.Move,
	values []string,
	text string,
) string {
	result := text

	for _, value := range values {
		moved, found := apply(moves, value)

		if !found {
			continue
		}

		for offset := 0; ; {
			i := strings.Index(result[offset:], value)

			if i == -1 {
				break
			}

			start := offset + i
			stop := start + len(value)
			before, _ := utf8.DecodeLastRuneInString(result[:start])
			after, _ := utf8.DecodeRuneInString(result[stop:])

			if (start > 0 && pointer.IsPathRune(before)) ||
				(stop < len(result) && pointer.IsPathRune(after)) {
				offset = stop

				continue
			}

			result = join.Empty(result[:start], moved, result[stop:])
			offset = start + len(moved)
		}
	}

	return result
}
