package sweep

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"os"
)

func synchronize(
	source string,
	destination string,
) constant.Action {
	di, e := os.Stat(destination)

	if e != nil {
		return constant.ActionCopy
	}

	si, f := os.Stat(source)

	if f != nil {
		return constant.ActionSkip
	}

	if si.Size() != di.Size() || si.ModTime().After(di.ModTime()) {
		return constant.ActionUpdate
	}

	return constant.ActionSkip
}
