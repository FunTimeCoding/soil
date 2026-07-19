package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/team"
	rawTeam "github.com/opsgenie/opsgenie-go-sdk-v2/team"
	"testing"
)

func TestTeam(t *testing.T) {
	assert.NotNil(t, team.New(&rawTeam.ListedTeams{}))
}
