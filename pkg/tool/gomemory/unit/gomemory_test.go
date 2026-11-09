package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemory/result"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"testing"
)

func TestProfileEmitsProfileAndPassesTopicAndDetail(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Profile = &client.ProfileResponse{Text: "alfa"}
	output, _, code := run(t, c, "profile", "--topic", "bravo", "--detail")
	assert.Integer(t, 0, code)
	var p client.ProfileResponse
	notation.MustDecode(output, &p, false)
	assert.String(t, "alfa", p.Text)
	assert.Count(t, 1, c.ProfileCalls)
	assert.FatalNotNil(t, c.ProfileCalls[0].Topic)
	assert.FatalNotNil(t, c.ProfileCalls[0].Detail)
	assert.String(t, "bravo", *c.ProfileCalls[0].Topic)
	assert.True(t, *c.ProfileCalls[0].Detail)
}

func TestProfileWithoutFlagsSendsNoTopicOrDetail(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	_, _, code := run(t, c, "profile")
	assert.Integer(t, 0, code)
	assert.Count(t, 1, c.ProfileCalls)
	assert.Nil(t, c.ProfileCalls[0].Topic)
	assert.Nil(t, c.ProfileCalls[0].Detail)
}

func TestProfilePrintsRenderedDocument(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Profile = &client.ProfileResponse{Text: "# always\n\nalfa\n"}
	output, failure, code := run(t, c, "profile", "--topic", "bravo")
	assert.Integer(t, 0, code)
	assert.String(t, "", failure)
	assert.String(t, "# always\n\nalfa\n", output)
}

func TestFailedProfileExitsWithItsMessage(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Failure = errors.New("foxtrot")
	output, failure, code := run(t, c, "profile")
	assert.Integer(t, 1, code)
	assert.String(t, "", output)
	assert.StringContains(t, "foxtrot", failure)
}

func TestStatisticEmitsSummaryForScope(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Tokens = &client.TokenSummary{Withheld: 3, Memories: 5}
	output, _, code := run(t, c, "statistic", "--scope", "kilo")
	assert.Integer(t, 0, code)
	var s client.TokenSummary
	notation.MustDecode(output, &s, false)
	assert.Integer(t, 3, s.Withheld)
	assert.Integer(t, 5, s.Memories)
	assert.Count(t, 1, c.TokenCalls)
	assert.FatalNotNil(t, c.TokenCalls[0].Scope)
	assert.String(t, "kilo", *c.TokenCalls[0].Scope)
}

func TestRelationsEmitsEdges(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Edges = []client.Relation{{SourceName: "alfa", TargetName: "bravo"}}
	output, _, code := run(
		t,
		c,
		"relations",
		"--type",
		"papa",
		"--scope",
		"kilo",
	)
	assert.Integer(t, 0, code)
	var r result.Relations
	notation.MustDecode(output, &r, false)
	assert.Count(t, 1, r.Relations)
	assert.String(t, "alfa", r.Relations[0].SourceName)
	assert.String(t, "bravo", r.Relations[0].TargetName)
	assert.FatalNotNil(t, r.Summary)
	assert.Integer(t, 1, r.Summary.Count)
	assert.Count(t, 1, c.RelationCalls)
	assert.FatalNotNil(t, c.RelationCalls[0].Type)
	assert.FatalNotNil(t, c.RelationCalls[0].Scope)
	assert.String(t, "papa", *c.RelationCalls[0].Type)
	assert.String(t, "kilo", *c.RelationCalls[0].Scope)
}

func TestRelationsUntypedAsksForUntypedFilter(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	_, _, code := run(t, c, "relations", "--untyped")
	assert.Integer(t, 0, code)
	assert.Count(t, 1, c.RelationCalls)
	assert.FatalNotNil(t, c.RelationCalls[0].Type)
	assert.String(t, "untyped", *c.RelationCalls[0].Type)
}

func TestFailedCallExitsWithItsMessage(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.Failure = errors.New("foxtrot")
	_, failure, code := run(t, c, "statistic")
	assert.Integer(t, 1, code)
	assert.StringContains(t, "foxtrot", failure)
}
