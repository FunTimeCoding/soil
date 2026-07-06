package goc

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"strings"
)

func configurationName(
	base string,
	selected string,
) string {
	var result string

	for _, f := range system.Files(join.Absolute(base, selected)) {
		if strings.HasSuffix(f, keyringSuffix) {
			return split.Dot(f)[2]
		}
	}

	return result
}
