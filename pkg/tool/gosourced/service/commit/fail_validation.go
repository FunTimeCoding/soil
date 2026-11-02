package commit

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func FailValidation(
	r *output.Results,
	message string,
) (*output.Results, error) {
	r.AddConcern(
		concern.NewFile(constant.ConcernValidation, message, "", false),
	)

	return r, nil
}
