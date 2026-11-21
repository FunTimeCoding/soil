package affected_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_row"
)

type Result struct {
	Identifier string
	Score      float64
	Severity   string
	Count      *affected_count.Count
	Rows       []*affected_row.Row
}
