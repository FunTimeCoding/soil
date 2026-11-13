package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"time"
)

func lastUsed(t time.Time) gomponents.Node {
	if t.IsZero() {
		return layout.OptionalTimeCell(nil, constant.CoverageNeverUsed)
	}

	return layout.TimeCell(t)
}
