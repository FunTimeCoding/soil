package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
	"strings"
)

func deadShort(candidate string) *pointer.Resolution {
	result := pointer.NewConcern(constant.DeadPointerKey)
	first, _, _ := strings.Cut(candidate, stringsConstant.Slash)

	if slices.Contains(constant.MediaTopLevelTypes, first) {
		result.Hint = join.Empty(
			constant.MediaHintPrefix,
			constant.SchemeMedia,
			candidate,
		)
	}

	return result
}
