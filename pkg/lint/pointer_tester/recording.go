package pointer_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func Recording(existing ...string) (lint.Checker, *[]string) {
	var seen []string

	return lint.Pointers(
		Resolver(existing...),
		func(
			path string,
			number int,
			span string,
			reason constant.Reason,
		) {
			seen = append(
				seen,
				fmt.Sprintf("%s:%d %s %s", path, number, reason, span),
			)
		},
	), &seen
}
