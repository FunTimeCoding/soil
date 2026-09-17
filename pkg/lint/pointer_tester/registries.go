package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func Registries(registries ...string) (lint.Checker, *[]string) {
	var seen []string
	r := pointer.New()
	r.Roots = Roots()
	r.Registries = registries

	return lint.Pointers(
		r,
		func(
			path string,
			number int,
			span string,
			reason constant.Reason,
		) {
			seen = append(seen, string(reason))
		},
	), &seen
}
