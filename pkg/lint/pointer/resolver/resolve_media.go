package resolver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"slices"
	"strings"
)

func resolveMedia(candidate string) *pointer.Resolution {
	top, subtype, _ := strings.Cut(
		strings.TrimPrefix(candidate, constant.SchemeMedia),
		stringsConstant.Slash,
	)

	if slices.Contains(constant.MediaTopLevelTypes, top) &&
		constant.MediaSubtype.MatchString(subtype) {
		return pointer.NewLive("")
	}

	return pointer.NewConcern(constant.UnregisteredMediaKey)
}
