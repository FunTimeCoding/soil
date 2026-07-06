package team_map

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	"testing"
)

func TestMap(t *testing.T) {
	assert.NotNil(t, New([]*team.Team{}))
}
