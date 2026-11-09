package pointer_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func Symbols(
	declared []string,
	types []string,
	tests []string,
	existing ...string,
) (lint.Checker, *[]string) {
	var seen []string
	r := Resolver(existing...)
	r.Declares = func(
		directory string,
		name string,
	) bool {
		return slices.Contains(
			declared,
			join.Empty(directory, stringsConstant.Dot, name),
		)
	}
	r.DeclaresType = func(
		directory string,
		name string,
	) bool {
		return slices.Contains(
			types,
			join.Empty(directory, stringsConstant.Dot, name),
		)
	}
	r.Tests = func(name string) bool {
		return slices.Contains(tests, name)
	}

	return lint.Pointers(
		r,
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
