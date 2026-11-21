package security_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	debianMock "github.com/funtimecoding/soil/pkg/vulnerability/debian/mock_client"
	nvdMock "github.com/funtimecoding/soil/pkg/vulnerability/national_database/mock_client"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{
		NationalDatabase: nvdMock.New(t),
		SecurityTracker:  debianMock.New(t),
		ExtendedTracker:  debianMock.New(t),
		Cache:            t.TempDir(),
		Reporter:         memory.New(),
		t:                t,
	}
}
