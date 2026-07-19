package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "MATTERMOST_HOST", constant.HostEnvironment)
	assert.String(t, "MATTERMOST_TOKEN", constant.TokenEnvironment)
	assert.String(t, "MATTERMOST_TEAM", constant.TeamEnvironment)
	assert.String(t, "MATTERMOST_CHANNEL", constant.ChannelEnvironment)
	assert.String(t, "construction", constant.Construction)
	assert.String(t, "hourglass_flowing_sand", constant.Hourglass)
	assert.String(t, "repeat", constant.Repeat)
	assert.String(t, "thread", constant.Thread)
}
