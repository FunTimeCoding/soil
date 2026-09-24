package crap

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/crap/option"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func RunAttribute(o *option.Attribute) int {
	if slices.Contains(o.Patterns, constant.AllPackages) {
		console.Format(constant.AttributeUnscoped)

		return 1
	}

	if v := index.ParallelTests(o.Root, o.Patterns...); len(v) > 0 {
		console.Format(constant.ParallelRefused, join.NewLine(v))

		return 1
	}

	i := index.Load(o.Root, o.Patterns...)
	m := Attribute(
		o.Root,
		i,
		coverage.TestPackages(o.Root, o.Patterns...),
		o.Patterns,
	)

	if o.Notation {
		console.Line(notation.MarshalIndent(m))

		return 0
	}

	printMatrix(m, i, o)

	return 0
}
