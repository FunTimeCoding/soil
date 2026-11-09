package model_context

import linkaceConstant "github.com/funtimecoding/soil/pkg/linkace/constant"

func offsetPage(
	offset int,
	limit int,
) int {
	if offset <= 0 {
		return 1
	}

	perPage := limit

	if perPage <= 0 {
		perPage = linkaceConstant.DefaultPerPage
	}

	return (offset / perPage) + 1
}
