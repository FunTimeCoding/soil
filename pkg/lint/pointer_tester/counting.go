package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/constant"
)

func Counting(existing ...string) (lint.Checker, *int) {
	count := 0

	return lint.Pointers(
		Resolver(existing...),
		func(string, int, string, constant.Reason) { count++ },
	), &count
}
