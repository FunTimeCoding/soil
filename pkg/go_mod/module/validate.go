package module

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/time"
)

func (m *Module) Validate() {
	if m.Release == nil {
		return
	}

	m.Months = time.MonthsBetween(*m.Release, m.now)

	if m.Months >= m.threshold {
		m.concern = append(m.concern, constant.ReleaseStale)
	}
}
