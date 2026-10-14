package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"testing"
	"time"
)

func createRecords(
	t *testing.T,
	s *store.Store,
	count int,
) {
	t.Helper()

	for i := range count {
		assert.FatalOnError(
			t,
			s.Create(
				record.Record{
					Fingerprint: fmt.Sprintf("fingerprint-%d", i),
					Name:        "HypervisorUnreachable",
					Severity:    "warning",
					Start:       time.Now().Add(-time.Hour),
				},
			),
		)
	}
}
