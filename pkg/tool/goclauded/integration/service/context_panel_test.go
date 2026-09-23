package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"testing"
)

func TestContextPanelScopeLineCountsLoadedAgainstCorpus(t *testing.T) {
	s := service_tester.NewPanel(t, "scope-session")
	s.Memory.Stats = &client.Statistics{
		Scopes: []client.NamedCount{{Count: 500}, {Name: "alfa", Count: 40}},
	}
	assert.String(
		t,
		"memory   10/500 loaded, alfa 40",
		s.Service.ContextPanel("scope-session"),
	)
}

func TestContextPanelTagLineCountsLoadedMembers(t *testing.T) {
	s := service_tester.NewPanel(t, "tag-session")
	s.Memory.Stats = corpus()
	s.Memory.Stats.Tags = []client.NamedCount{
		{Name: "bravo", Count: 4, Identifiers: new([]int64{102, 103, 901})},
		{Name: "charlie", Count: 7},
	}
	assert.String(
		t,
		"memory   10/500 loaded\ntags     bravo 2/4",
		s.Service.ContextPanel("tag-session"),
	)
}

func TestContextPanelOmitsTagLineWithoutMembership(t *testing.T) {
	s := service_tester.NewPanel(t, "bare-tag-session")
	s.Memory.Stats = corpus()
	s.Memory.Stats.Tags = []client.NamedCount{{Name: "charlie", Count: 7}}
	assert.String(
		t,
		"memory   10/500 loaded",
		s.Service.ContextPanel("bare-tag-session"),
	)
}

func TestContextPanelDoorsReachUnloadedNeighbors(t *testing.T) {
	s := service_tester.NewPanel(t, "door-session")
	s.Memory.Stats = corpus()
	s.Memory.Edges = []client.Relation{
		fixture.Edge(109, "kilo", 301, "november"),
	}
	assert.String(
		t,
		"memory   10/500 loaded\ndoor     301  november   informs <- kilo",
		s.Service.ContextPanel("door-session"),
	)
}

func TestContextPanelDoorsNeedExactlyOneLoadedEnd(t *testing.T) {
	s := service_tester.NewPanel(t, "closed-session")
	s.Memory.Stats = corpus()
	s.Memory.Edges = []client.Relation{
		fixture.Edge(109, "kilo", 110, "lima"),
		fixture.Edge(301, "november", 302, "oscar"),
	}
	assert.String(
		t,
		"memory   10/500 loaded",
		s.Service.ContextPanel("closed-session"),
	)
}

func TestContextPanelLabelsRelationsWithoutType(t *testing.T) {
	s := service_tester.NewPanel(t, "untyped-session")
	s.Memory.Stats = corpus()
	s.Memory.Edges = []client.Relation{
		{
			SourceIdentifier: 109,
			SourceName:       "kilo",
			TargetIdentifier: 301,
			TargetName:       "november",
		},
	}
	assert.String(
		t,
		"memory   10/500 loaded\ndoor     301  november   untyped <- kilo",
		s.Service.ContextPanel("untyped-session"),
	)
}

func TestContextPanelKeepsFirstDoorToRepeatedNeighbor(t *testing.T) {
	s := service_tester.NewPanel(t, "repeat-session")
	s.Memory.Stats = corpus()
	s.Memory.Edges = []client.Relation{
		fixture.Edge(109, "kilo", 301, "november"),
		fixture.Edge(110, "lima", 301, "november"),
	}
	assert.String(
		t,
		"memory   10/500 loaded\ndoor     301  november   informs <- kilo",
		s.Service.ContextPanel("repeat-session"),
	)
}

func TestContextPanelSortsDoorsByIdentifier(t *testing.T) {
	s := service_tester.NewPanel(t, "sorted-session")
	s.Memory.Stats = corpus()
	s.Memory.Edges = []client.Relation{
		fixture.Edge(109, "kilo", 302, "oscar"),
		fixture.Edge(110, "lima", 301, "november"),
	}
	assert.String(
		t,
		"memory   10/500 loaded\ndoor     301  november   informs <- lima\ndoor     302  oscar   informs <- kilo",
		s.Service.ContextPanel("sorted-session"),
	)
}

func TestContextPanelEmptyWithoutStatistics(t *testing.T) {
	s := service_tester.NewPanel(t, "statless-session")
	assert.String(t, "", s.Service.ContextPanel("statless-session"))
}
