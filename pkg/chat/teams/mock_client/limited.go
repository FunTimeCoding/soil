package mock_client

import "github.com/funtimecoding/soil/pkg/chat/constant"

func limited[T any](
	v []T,
	limit int,
) []T {
	if limit <= 0 {
		limit = constant.TeamsDefaultLimit
	}

	return append([]T{}, v[:min(limit, len(v))]...)
}
