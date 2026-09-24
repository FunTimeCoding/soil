package measure

import (
	"github.com/funtimecoding/soil/pkg/console/table"
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/measure/summary"
)

func addSummary(
	t *table.Table,
	s *summary.Summary,
) {
	t.Add(
		s.Language,
		integers.ToString(s.Files),
		integers.ToString(s.Count.Blank),
		integers.ToString(s.Count.Comment),
		integers.ToString(s.Count.Code),
	)
}
