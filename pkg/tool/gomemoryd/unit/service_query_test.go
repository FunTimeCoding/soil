package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func TestSearchHitResolvesToItsMemory(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	found, f := o.Service.SearchRelevant(constant.FixtureContent, 10, nil)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, found)
	assert.Integer(t, int(m.Identifier), int(found[0].Identifier))
	assert.String(t, "retry policy", found[0].Name)
}

func TestSearchHitWithoutMemoryIsDropped(t *testing.T) {
	o := service_tester.New(t)
	_, e := o.Indexer.PostDocument(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "memories",
			Path:       "memory/99",
			Body:       constant.FixtureContent,
		},
	)
	assert.FatalOnError(t, e)
	found, f := o.Service.SearchRelevant(constant.FixtureContent, 10, nil)
	assert.FatalOnError(t, f)
	assert.Count(t, 0, found)
}

func TestSearchGoquerydRefusesKeepsItsClass(t *testing.T) {
	o := service_tester.New(t)
	o.Indexer.Failure = unavailable.Format("embedder unreachable")
	_, e := o.Service.SearchRelevant("retry", 10, nil)
	assert.True(t, unavailable.Is(e))
}

func TestCompletionsKeepLatestPerSession(t *testing.T) {
	o := service_tester.New(t)
	postCompletion(t, o, "alfa/2", "2026-11-09T12:00:00+01:00")
	postCompletion(t, o, "alfa/1", "2026-11-09T10:00:00+01:00")
	postCompletion(t, o, "bravo/1", "2026-11-09T10:30:00Z")
	postCompletion(t, o, "alfa/3", "2026-11-09T09:00:00+01:00")
	found, e := o.Service.ListCompletions()
	assert.FatalOnError(t, e)
	assert.Count(t, 2, found)
	assert.String(t, "alfa/2", found[0].Path)
	assert.String(t, "alfa/2", *found[0].Body)
	assert.String(t, "bravo/1", found[1].Path)
}

func TestCompletionsOrderByCreationAcrossOffsets(t *testing.T) {
	o := service_tester.New(t)
	postCompletion(t, o, "alfa/1", "2026-11-09T11:00:00+01:00")
	postCompletion(t, o, "bravo/1", "2026-11-09T10:30:00Z")
	found, e := o.Service.ListCompletions()
	assert.FatalOnError(t, e)
	assert.Count(t, 2, found)
	assert.String(t, "bravo/1", found[0].Path)
}

func TestCompletionsStopAtFiveSessions(t *testing.T) {
	o := service_tester.New(t)
	postCompletion(t, o, "alfa/1", "2026-11-09T10:00:00+01:00")
	postCompletion(t, o, "bravo/1", "2026-11-09T11:00:00+01:00")
	postCompletion(t, o, "charlie/1", "2026-11-09T12:00:00+01:00")
	postCompletion(t, o, "delta/1", "2026-11-09T13:00:00+01:00")
	postCompletion(t, o, "echo/1", "2026-11-09T14:00:00+01:00")
	postCompletion(t, o, "foxtrot/1", "2026-11-09T15:00:00+01:00")
	found, e := o.Service.ListCompletions()
	assert.FatalOnError(t, e)
	assert.Count(t, 5, found)
	assert.String(t, "foxtrot/1", found[0].Path)
	assert.String(t, "bravo/1", found[4].Path)
}

func TestCompletionWithoutTimestampSortsLast(t *testing.T) {
	o := service_tester.New(t)
	postCompletion(t, o, "alfa/1", "")
	postCompletion(t, o, "bravo/1", "2026-11-09T10:00:00+01:00")
	found, e := o.Service.ListCompletions()
	assert.FatalOnError(t, e)
	assert.Count(t, 2, found)
	assert.String(t, "alfa/1", found[1].Path)
}
