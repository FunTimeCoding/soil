package table

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (t *Table) Numeric(columns ...int) *Table {
	for _, c := range columns {
		if c < 0 || c >= constant.NumericColumnLimit {
			errors.PanicOnError(
				errors.Format("numeric column outside the shared rules", c),
			)
		}

		t.numeric[c] = true
	}

	return t
}
