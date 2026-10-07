package message_meta

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"

func New(
	typeValue string,
	statistic response.Statistic,
) *Meta {
	return &Meta{Type: typeValue, Statistic: statistic}
}
