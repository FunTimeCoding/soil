package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func failValidation(
	r *output.Results,
	message string,
) (*output.Results, error) {
	r.AddConcern(concern.NewFile("validation", message, "", false))

	return r, nil
}
