package query_result

import "github.com/prometheus/common/model"

type Result struct {
	Value    model.Value
	Warnings []string
	Notices  []string
}
