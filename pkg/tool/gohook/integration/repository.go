package integration

import (
	"github.com/funtimecoding/soil/pkg/git/unit/repository_tester"
	"testing"
)

func repository(t *testing.T) string {
	t.Helper()

	return repository_tester.New(t).Clone
}
