package rotation

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"testing"
)

func TestRotation(t *testing.T) {
	assert.NotNil(t, New(&schedule.Rotation{}))
}
