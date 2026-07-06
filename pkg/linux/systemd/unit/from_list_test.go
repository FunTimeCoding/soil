package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFromList(t *testing.T) {
	assert.Any(
		t,
		&Unit{
			Name:        "smartd.service",
			Load:        "loaded",
			Active:      "failed",
			Sub:         "failed",
			Description: "Self Monitoring and Reporting Technology (SMART) Daemon",
		},
		FromList(
			"smartd.service loaded failed failed Self Monitoring and Reporting Technology (SMART) Daemon",
		),
	)
}
