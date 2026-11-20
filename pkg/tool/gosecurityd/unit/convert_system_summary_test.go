package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"testing"
	"time"
)

func TestSummaryCarriesAges(t *testing.T) {
	verified := time.Date(2026, time.November, 19, 8, 0, 0, 0, time.UTC)
	v := system_result.NewSummary(system.New("tango", "trixie", 0))
	v.Stale = true
	v.TrackerStale = true
	v.TrackerVerified = verified
	s := convert.SystemSummary(v)
	assert.True(t, s.Stale)
	assert.True(t, s.TrackerStale)
	assert.NotNil(t, s.TrackerVerified)
	assert.Time(t, verified, *s.TrackerVerified)
	assert.Nil(
		t,
		convert.SystemSummary(
			system_result.NewSummary(system.New("uniform", "trixie", 0)),
		).TrackerVerified,
	)
}
