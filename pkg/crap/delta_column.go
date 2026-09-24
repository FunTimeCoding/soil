package crap

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/entry"
)

func deltaColumn(
	e *entry.Entry,
	tolerance float64,
) string {
	switch {
	case e.IsNew():
		return constant.DeltaNew
	case e.Unchanged(tolerance):
		return constant.DeltaSame
	default:
		return fmt.Sprintf(constant.DeltaFormat, e.Delta())
	}
}
