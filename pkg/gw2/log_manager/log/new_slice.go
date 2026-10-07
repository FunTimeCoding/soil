package log

import "github.com/funtimecoding/soil/pkg/gw2/log_manager/response"

func NewSlice(l []*response.Log) []*Log {
	var result []*Log

	for _, o := range l {
		result = append(result, New(o))
	}

	SortByTime(result)

	return result
}
