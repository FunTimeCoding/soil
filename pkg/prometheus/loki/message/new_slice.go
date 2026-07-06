package message

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"

func NewSlice(v *query_result.Result) ([]*Message, *Meta) {
	var result []*Message

	for _, e := range v.Result {
		for _, a := range e.Values {
			result = append(result, New(a, &e.Stream))
		}
	}

	return result, &Meta{Type: v.ResultType, Statistic: v.Stats}
}
