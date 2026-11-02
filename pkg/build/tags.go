package build

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Tags(v string) string {
	if v == "" {
		return soil.TimeZoneTag
	}

	return join.Comma([]string{soil.TimeZoneTag, v})
}
