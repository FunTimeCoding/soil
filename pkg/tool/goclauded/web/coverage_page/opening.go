package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
)

func opening(e *coverage.Server) *opening_row.OpeningRow {
	return opening_row.New(
		component.RowMark(constant.CoverageRowPrefix, e.Name),
		component.RowPath(
			constant.CoveragePath,
			constant.OpenAction,
			constant.CoverageServer,
			e.Name,
		),
		component.RowPath(
			constant.CoveragePath,
			constant.CloseAction,
			constant.CoverageServer,
			e.Name,
		),
	).Cells(cells(e)...)
}
