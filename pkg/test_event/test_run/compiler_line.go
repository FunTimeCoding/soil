package test_run

import (
	"github.com/funtimecoding/soil/pkg/test_event/constant"
	"strings"
)

func (r *Run) CompilerLine() string {
	output := r.BuildOutput()

	for _, l := range output {
		if !strings.HasPrefix(l, constant.PackageHeader) {
			return l
		}
	}

	if len(output) == 0 {
		return ""
	}

	return output[0]
}
