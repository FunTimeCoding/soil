package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"path/filepath"
	"strings"
	"testing"
)

func TestEnrichSessionIndexesSearch(t *testing.T) {
	s := service_tester.New(t)
	s.WriteSessionFile("alfa", "alfa-slug")
	s.Transcript.EnrichSession("alfa")
	r, e := s.Index.Search("login bug", nil, 20)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(r))
	assert.String(t, "alfa", r[0].Session)
}

func TestCatchUpSearchIndexesAndDropsVanished(t *testing.T) {
	s := service_tester.New(t)
	s.WriteSessionFile("alfa", "alfa-slug")
	s.WriteSessionFile("bravo", "bravo-slug")
	assert.Count(t, 0, s.Search.CatchUpSearch(t.Context()))
	r, e := s.Index.Search("login bug", nil, 20)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, len(r))
	system.Remove(filepath.Join(s.Harbor, "bravo.jsonl"))
	assert.Count(t, 0, s.Search.CatchUpSearch(t.Context()))
	r, e = s.Index.Search("login bug", nil, 20)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(r))
	assert.String(t, "alfa", r[0].Session)
}

func TestReadConversationTruncatesLongBlocks(t *testing.T) {
	s := service_tester.New(t)
	system.WriteFile(
		filepath.Join(s.Harbor, "long.jsonl"),
		[]byte(
			fmt.Sprintf(
				`{"uuid":"w1","type":"user","timestamp":"2026-10-08T01:00:00Z","message":{"role":"user","content":"%s"}}%s`,
				strings.Repeat("x", 2005),
				"\n",
			),
		),
		0644,
	)
	assert.Count(t, 0, s.Search.CatchUpSearch(t.Context()))
	blocks, e := s.Search.ReadConversation("long", "w1", 0)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(blocks))
	assert.String(
		t,
		join.Empty(strings.Repeat("x", 2000), "\n(5 more characters)"),
		blocks[0].Text,
	)
}

func TestDeleteSessionReceiptCountsSearchEntries(t *testing.T) {
	s := service_tester.New(t)
	s.WriteSessionFile("doomed", "doomed-slug")
	s.Transcript.PopulateCache()
	s.Transcript.CheckConsistency()
	s.Transcript.EnrichSession("doomed")
	r, e := s.Session.DeleteSession("doomed", s.Session.DeleteHash("doomed"))
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, int(r.SearchEntries))
	found, e := s.Index.Search("login bug", nil, 20)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, len(found))
}
