package message

import (
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
	"github.com/funtimecoding/soil/pkg/prometheus/types/message_meta"
)

func NewSlice(v *query_result.Result) ([]*Message, *message_meta.Meta) {
	var result []*Message

	for _, e := range v.Result {
		for _, a := range e.Values {
			result = append(result, New(a, &e.Stream))
		}
	}

	return result, message_meta.New(v.ResultType, v.Stats)
}
