package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func TestASearchHitResolvesToItsMemory(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(scopedOption("retry policy", ""))
	assert.FatalOnError(t, e)
	found, f := o.Service.SearchRelevant(constant.FixtureContent, 10, nil)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, found)
	assert.Integer(t, int(m.Identifier), int(found[0].Identifier))
	assert.String(t, "retry policy", found[0].Name)
}

func TestASearchHitWithoutAMemoryIsDropped(t *testing.T) {
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

func TestASearchGoquerydRefusesKeepsItsClass(t *testing.T) {
	o := service_tester.New(t)
	o.Indexer.Failure = unavailable.Format("embedder unreachable")
	_, e := o.Service.SearchRelevant("retry", 10, nil)
	assert.True(t, unavailable.Is(e))
}

func TestCompletionsStopAtFiveSessions(t *testing.T) {
	o := service_tester.New(t)
	metadata := map[string][]string{"source_type": {"session-completion"}}

	for _, path := range []string{
		"alfa/1", "alfa/2", "bravo/1", "charlie/1",
		"delta/1", "echo/1", "foxtrot/1",
	} {
		_, e := o.Indexer.PostDocument(
			t.Context(),
			client.PostDocumentJSONRequestBody{
				Collection: "completions",
				Path:       path,
				Body:       path,
				Metadata:   &metadata,
			},
		)
		assert.FatalOnError(t, e)
	}

	found, f := o.Service.ListCompletions()
	assert.FatalOnError(t, f)
	assert.Count(t, 6, found)
	assert.String(t, "echo/1", found[5].Path)
}
