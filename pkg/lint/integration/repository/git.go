package repository

import (
	"github.com/funtimecoding/soil/pkg/system/run"
	"slices"
	"testing"
)

func git(
	t *testing.T,
	directory string,
	arguments ...string,
) {
	t.Helper()
	r := run.New()
	r.Directory = directory
	r.Start(
		slices.Concat(
			[]string{
				"git",
				"-c",
				"user.name=alfa",
				"-c",
				"user.email=alfa@example.org",
			},
			arguments,
		)...,
	)
}
