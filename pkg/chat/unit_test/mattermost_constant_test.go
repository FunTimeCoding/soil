package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"testing"
)

func TestClient(t *testing.T) {
	assert.String(t, "MATTERMOST_HOST", constant.MattermostHostEnvironment)
	assert.String(t, "MATTERMOST_TOKEN", constant.MattermostTokenEnvironment)
	assert.String(t, "MATTERMOST_TEAM", constant.MattermostTeamEnvironment)
	assert.String(
		t,
		"MATTERMOST_CHANNEL",
		constant.MattermostChannelEnvironment,
	)
	assert.String(
		t,
		"MATTERMOST_INSECURE",
		constant.MattermostInsecureEnvironment,
	)
	assert.String(t, "construction", constant.MattermostConstruction)
	assert.String(t, "hourglass_flowing_sand", constant.MattermostHourglass)
	assert.String(t, "repeat", constant.MattermostRepeat)
	assert.String(t, "thread", constant.MattermostThread)
}
