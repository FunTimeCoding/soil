package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/rotation"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"testing"
)

func TestRotation(t *testing.T) {
	assert.NotNil(t, rotation.New(&schedule.Rotation{}))
}
