package measure

import (
	"github.com/funtimecoding/soil/pkg/console/table"
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/measure/constant"
	"github.com/funtimecoding/soil/pkg/measure/result"
)

func printFiles(r *result.Result) {
	t := table.New(
		constant.ColumnLanguage,
		constant.ColumnBlank,
		constant.ColumnComment,
		constant.ColumnCode,
		constant.ColumnPath,
	).Right(1, 2, 3)

	for _, f := range r.ByPath() {
		t.Add(
			f.Language,
			integers.ToString(f.Count.Blank),
			integers.ToString(f.Count.Comment),
			integers.ToString(f.Count.Code),
			f.Path,
		)
	}

	t.Print()
}
