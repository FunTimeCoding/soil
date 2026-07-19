package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/override"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"testing"
)

func TestOverride(t *testing.T) {
	assert.NotNil(t, override.New(&schedule.ScheduleOverride{}))
}
