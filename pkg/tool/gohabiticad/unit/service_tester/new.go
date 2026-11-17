package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{
		Source:   mock_client.New(t),
		Reporter: memory.New(),
		t:        t,
	}
}
