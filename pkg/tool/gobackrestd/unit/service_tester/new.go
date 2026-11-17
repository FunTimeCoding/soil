package service_tester

import (
	"github.com/funtimecoding/soil/pkg/backrest/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{Source: mock_client.New(), Reporter: memory.New(), t: t}
}
