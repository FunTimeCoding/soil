package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "MATTERMOST_HOST", HostEnvironment)
	assert.String(t, "MATTERMOST_TOKEN", TokenEnvironment)
	assert.String(t, "MATTERMOST_TEAM", TeamEnvironment)
	assert.String(t, "MATTERMOST_CHANNEL", ChannelEnvironment)
	assert.String(t, "construction", Construction)
	assert.String(t, "hourglass_flowing_sand", Hourglass)
	assert.String(t, "repeat", Repeat)
	assert.String(t, "thread", Thread)
}
