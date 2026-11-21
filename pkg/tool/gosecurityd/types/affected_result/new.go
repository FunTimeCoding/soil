package affected_result

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_row"
)

func New(identifier string) *Result {
	return &Result{
		Identifier: identifier,
		Count:      affected_count.New(),
		Rows:       []*affected_row.Row{},
	}
}
