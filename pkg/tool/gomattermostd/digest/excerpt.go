package digest

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"strings"
)

func excerpt(text string) string {
	flat := strings.Join(strings.Fields(text), " ")
	r := []rune(flat)

	if len(r) <= constant.ExcerptLength {
		return flat
	}

	return join.Empty(
		string(r[:constant.ExcerptLength]),
		constant.TruncationMarker,
	)
}
