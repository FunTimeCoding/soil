package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"testing"
	"time"
)

func TestACruiseNeverWokenShowsOnlyItsSetting(t *testing.T) {
	assert.String(
		t,
		"cruise paced every 15m",
		goclaude.CruiseSummary(
			&client.CruiseResponse{Mode: "paced", PaceMinutes: new(15)},
		),
	)
}

func TestAWokenCruiseCountsItsWakes(t *testing.T) {
	assert.String(
		t,
		"cruise on, woken 3 times, last 2026-11-02T01:30:00Z",
		goclaude.CruiseSummary(
			&client.CruiseResponse{
				Mode:       "on",
				Woken:      3,
				PromotedAt: new(time.Date(2026, 11, 2, 1, 30, 0, 0, time.UTC)),
			},
		),
	)
}
