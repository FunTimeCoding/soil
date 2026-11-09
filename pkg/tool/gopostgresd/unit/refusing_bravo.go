package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"testing"
)

func refusingBravo(t *testing.T) *service_tester.Tester {
	t.Helper()
	s, clients := service_tester.New(t, "alfa", "bravo")
	clients["bravo"].Failure = validation.New("bad statement")

	return s
}
