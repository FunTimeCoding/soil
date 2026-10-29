package view

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/face"
)

func report(
	r face.Reporter,
	caught any,
) (string, string) {
	if e, okay := caught.(error); okay && !classify.Reportable(e) {
		return e.Error(), ""
	}

	return fmt.Sprintf("%v", caught), r.Recover(caught)
}
