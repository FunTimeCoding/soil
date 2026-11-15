package client_tester

import (
	"github.com/funtimecoding/soil/pkg/github"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *github.Client {
	t.Helper()
	result := github.NewServer("token", address)
	t.Cleanup(result.Close)

	return result
}
