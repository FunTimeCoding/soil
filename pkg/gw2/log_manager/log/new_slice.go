package log

import "github.com/funtimecoding/soil/pkg/gw2/log_manager"

func NewSlice(l []*log_manager.Log) []*Log {
	var result []*Log

	for _, o := range l {
		result = append(result, New(o))
	}

	SortByTime(result)

	return result
}
