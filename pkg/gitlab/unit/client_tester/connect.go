package client_tester

import (
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
)

func Connect(
	t *testing.T,
	address string,
	o ...gitlab.Option,
) *gitlab.Client {
	t.Helper()
	result := gitlab.New(
		web.TrimScheme(address),
		"test-token",
		append(
			[]gitlab.Option{gitlab.WithInsecure(), gitlab.WithRetryMax(0)},
			o...,
		)...,
	)
	t.Cleanup(result.Close)

	return result
}
