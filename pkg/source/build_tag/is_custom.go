package build_tag

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"strings"
)

func isCustom(tag string) bool {
	if constant.KnownTags[tag] {
		return false
	}

	if strings.HasPrefix(tag, "go1.") {
		return false
	}

	if strings.HasPrefix(tag, "goexperiment.") {
		return false
	}

	return true
}
