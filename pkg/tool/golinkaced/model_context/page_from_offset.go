package model_context

import "github.com/funtimecoding/soil/pkg/linkace/constant"

func pageFromOffset(
	offset int,
	limit int,
) int {
	if offset <= 0 {
		return 1
	}

	perPage := limit

	if perPage <= 0 {
		perPage = constant.DefaultPerPage
	}

	return (offset / perPage) + 1
}
