package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team_map"
	"testing"
)

func TestOpsgenieTeamMap(t *testing.T) {
	assert.NotNil(t, team_map.New([]*team.Team{}))
}
