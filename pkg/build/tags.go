package build

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Tags(v string) string {
	if v == "" {
		return constant.TimeZoneTag
	}

	return join.Comma([]string{constant.TimeZoneTag, v})
}
