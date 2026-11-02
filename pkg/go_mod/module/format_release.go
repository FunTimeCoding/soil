package module

import (
	"fmt"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (m *Module) formatRelease() string {
	if m.Release == nil {
		return ""
	}

	return fmt.Sprintf(
		"%s (%d months)",
		m.Release.Format(timeConstant.DateYear),
		m.Months,
	)
}
