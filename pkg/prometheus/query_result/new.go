package query_result

import "github.com/prometheus/common/model"

func New(
	v model.Value,
	w []string,
) *Result {
	return &Result{Value: v, Warnings: w}
}
