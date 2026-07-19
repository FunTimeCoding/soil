package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/linux/systemd/unit"
	"testing"
)

func TestFromList(t *testing.T) {
	assert.Any(
		t,
		&unit.Unit{
			Name:        "smartd.service",
			Load:        "loaded",
			Active:      "failed",
			Sub:         "failed",
			Description: "Self Monitoring and Reporting Technology (SMART) Daemon",
		},
		unit.FromList(
			"smartd.service loaded failed failed Self Monitoring and Reporting Technology (SMART) Daemon",
		),
	)
}
