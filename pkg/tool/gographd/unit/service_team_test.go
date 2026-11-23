package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/teams/team"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/service_tester"
	"testing"
)

func TestChannelsByTeamNameIgnoreCase(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	channels, e := o.Service().Channels("security")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, channels)
	assert.String(t, "channel-1", channels[0].Identifier)
}

func TestChannelsByTeamIdentifier(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	channels, e := o.Service().Channels("team-1")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, channels)
}

func TestUnknownTeamNotFound(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	_, e := o.Service().Channels("Finance")
	assert.True(t, not_found.Is(e))
	assert.String(t, "team not found: Finance", e.Error())
}

func TestEmptyTeamInvalid(t *testing.T) {
	_, e := service_tester.New(t).Service().Channels("")
	assert.True(t, validation.Is(e))
}

func TestSharedTeamNameAmbiguous(t *testing.T) {
	o := service_tester.New(t)
	o.Seed()
	o.Source.AddTeam(team.New("team-3", "Security"))
	_, e := o.Service().Channels("Security")
	assert.True(t, ambiguous.Is(e))
}
