package measure

import (
	"github.com/funtimecoding/soil/pkg/console/table"
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/result"
)

func printLanguages(r *result.Result) {
	t := table.New(
		constant.ColumnLanguage,
		constant.ColumnFiles,
		constant.ColumnBlank,
		constant.ColumnComment,
		constant.ColumnCode,
	).Right(1, 2, 3, 4)

	for _, s := range r.ByLanguage() {
		addSummary(t, s)
	}

	addSummary(t, r.Total())
	t.Print()
}
