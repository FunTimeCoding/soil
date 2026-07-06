package convert

import "github.com/funtimecoding/soil/pkg/generative/model_context/paginate"

func NewSilenceResult(
	v []*SlimSilence,
	limit int,
	offset int,
) *SlimSilenceResult {
	return &SlimSilenceResult{
		Silences: paginate.Slice(v, limit, offset),
		Total:    len(v),
	}
}
