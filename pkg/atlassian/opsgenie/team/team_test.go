package team

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
	"testing"
)

func TestTeam(t *testing.T) {
	assert.NotNil(t, New(&team.ListedTeams{}))
}
