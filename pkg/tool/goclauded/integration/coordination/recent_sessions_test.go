package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func targetByIdentifier(
	targets []*connector.Target,
	identifier string,
) *connector.Target {
	for _, t := range targets {
		if t.Identifier == identifier {
			return t
		}
	}

	return nil
}

func TestRecentSessionsCarriesLabels(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	set, e := a.RestClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		labelBody("environment", "staging", "reconciler"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, set.StatusCode())
	targets, f := s.Connector(t).RecentSessions(25)
	assert.FatalOnError(t, f)
	assert.True(t, len(targets) > 0)
	found := targetByIdentifier(targets, a.UUID)
	assert.NotNil(t, found)
	assert.String(t, "staging", found.Labels["environment"])
}

func TestRecentSessionsKeepsUnlabelledSessions(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	targets, e := s.Connector(t).RecentSessions(25)
	assert.FatalOnError(t, e)
	found := targetByIdentifier(targets, a.UUID)
	assert.NotNil(t, found)
	assert.Integer(t, 0, len(found.Labels))
}
