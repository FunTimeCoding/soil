package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/schedule"
	rawSchedule "github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"testing"
)

func TestSchedule(t *testing.T) {
	assert.NotNil(t, schedule.New(&rawSchedule.Schedule{}, nil))
}
