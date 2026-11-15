package gobreak

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
)

func expectProblem(f *break_file.File) string {
	for _, b := range f.Breaks {
		switch b.Expect {
		case "", constant.StatusKilled, constant.StatusSurvived:
		default:
			return fmt.Sprintf(constant.UnknownExpect, b.Name, b.Expect)
		}
	}

	return ""
}
