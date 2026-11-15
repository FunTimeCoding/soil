package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/search_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestSearchIndexMessages(t *testing.T) {
	x, _ := search_tester.New(t)
	r := search_tester.Search(t, x, "legacy")
	assert.Integer(t, 1, len(r))
	assert.String(t, "s1", r[0].Session)
	assert.Integer(t, 2, r[0].Count)
	assert.String(t, "2026-10-07T20:01:00Z", r[0].Latest)
	assert.Integer(t, 2, len(r[0].Hits))
	assert.String(t, "a1", r[0].Hits[0].Identifier)
	assert.String(t, "message", r[0].Hits[0].Kind)
	assert.String(
		t,
		"advertise legacy versions through the discover RPC",
		r[0].Hits[0].Snippet,
	)
	assert.String(t, "u1", r[0].Hits[1].Identifier)
}

func TestSearchIndexKinds(t *testing.T) {
	x, _ := search_tester.New(t)
	assert.Integer(t, 0, len(search_tester.Search(t, x, "StreamableHTTP")))
	r := search_tester.Search(t, x, "StreamableHTTP", constant.BlockEdit)
	assert.Integer(t, 1, len(r))
	assert.String(t, "a2", r[0].Hits[0].Identifier)
	assert.String(t, "edit", r[0].Hits[0].Kind)
	assert.String(
		t,
		"Edit /tmp/setup.go WithStreamableHTTPProtocolVersions heartbeat",
		r[0].Hits[0].Snippet,
	)
}

func TestSearchIndexCallText(t *testing.T) {
	x, _ := search_tester.New(t)
	r := search_tester.Search(t, x, "run tests", constant.BlockCall)
	assert.Integer(t, 1, len(r))
	assert.String(t, "a3", r[0].Hits[0].Identifier)
	assert.String(t, "Bash go test ./... run the tests", r[0].Hits[0].Snippet)
}

func TestSearchIndexConversationScope(t *testing.T) {
	x, _ := search_tester.New(t)
	r := search_tester.Search(t, x, "pin discover")
	assert.Integer(t, 1, len(r))
	assert.Integer(t, 2, r[0].Count)
	assert.String(t, "a1", r[0].Hits[0].Identifier)
	assert.String(t, "u1", r[0].Hits[1].Identifier)
	assert.Integer(t, 0, len(search_tester.Search(t, x, "legacy nonexistent")))
}

func TestSearchIndexAllTermsFirst(t *testing.T) {
	x, _ := search_tester.New(t)
	r := search_tester.Search(t, x, "legacy pin")
	assert.Integer(t, 1, len(r))
	assert.String(t, "u1", r[0].Hits[0].Identifier)
	assert.String(t, "a1", r[0].Hits[1].Identifier)
}

func TestSearchIndexShortTerm(t *testing.T) {
	x, _ := search_tester.New(t)
	r := search_tester.Search(t, x, "te", constant.BlockCall)
	assert.Integer(t, 1, len(r))
	assert.String(t, "a3", r[0].Hits[0].Identifier)
}

func TestSearchIndexSkipsResultsAndMeta(t *testing.T) {
	x, _ := search_tester.New(t)
	kinds := []string{
		constant.BlockMessage,
		constant.BlockEdit,
		constant.BlockCall,
	}
	assert.Integer(t, 0, len(search_tester.Search(t, x, "PASS", kinds...)))
	assert.Integer(t, 0, len(search_tester.Search(t, x, "clear", kinds...)))
	assert.Integer(t, 0, len(search_tester.Search(t, x, "Caveat", kinds...)))
}

func TestSearchIndexAppendCompleteLinesOnly(t *testing.T) {
	x, path := search_tester.New(t)
	full := search_tester.Line(
		"a9",
		"assistant",
		"2026-10-07T21:00:00Z",
		false,
		`[{"type":"text","text":"appended legacy note"}]`,
	)
	f, e := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	assert.FatalOnError(t, e)
	_, e = f.WriteString(full[:20])
	assert.FatalOnError(t, e)
	x.Append("s1", path)
	assert.Integer(t, 0, len(search_tester.Search(t, x, "appended")))
	_, e = f.WriteString(full[20:])
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, f.Close())
	x.Append("s1", path)
	assert.Integer(t, 1, len(search_tester.Search(t, x, "appended")))
	assert.Integer(t, 3, search_tester.Search(t, x, "legacy")[0].Count)
}

func TestSearchIndexDeleteSession(t *testing.T) {
	x, _ := search_tester.New(t)
	count, e := x.DeleteSession("s1")
	assert.FatalOnError(t, e)
	assert.Integer(t, 4, count)
	assert.Integer(t, 0, len(search_tester.Search(t, x, "legacy")))
}

func TestSearchIndexReindex(t *testing.T) {
	x, path := search_tester.New(t)
	system.WriteFile(
		path,
		[]byte(
			search_tester.Line(
				"n1",
				"user",
				"2026-10-08T09:00:00Z",
				false,
				`"a rewritten transcript"`,
			),
		),
		0644,
	)
	x.Reindex("s1", path)
	assert.Integer(t, 0, len(search_tester.Search(t, x, "legacy")))
	assert.Integer(t, 1, len(search_tester.Search(t, x, "rewritten")))
}

func TestSearchIndexShrunkFileReindexed(t *testing.T) {
	x, path := search_tester.New(t)
	system.WriteFile(
		path,
		[]byte(
			search_tester.Line(
				"n1",
				"user",
				"2026-10-08T09:00:00Z",
				false,
				`"a shorter transcript"`,
			),
		),
		0644,
	)
	x.Append("s1", path)
	assert.Integer(t, 0, len(search_tester.Search(t, x, "legacy")))
	assert.Integer(t, 1, len(search_tester.Search(t, x, "shorter")))
}

func TestSearchIndexMissingFile(t *testing.T) {
	x, path := search_tester.New(t)
	x.Append("s9", filepath.Join(filepath.Dir(path), "absent.jsonl"))
	sessions, e := x.Sessions()
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"s1"}, sessions)
}

func TestSearchIndexRole(t *testing.T) {
	x, _ := search_tester.New(t)
	r := search_tester.Search(t, x, "legacy")
	assert.String(t, "assistant", r[0].Hits[0].Role)
	assert.String(t, "user", r[0].Hits[1].Role)
}

func TestSearchIndexWindow(t *testing.T) {
	x, _ := search_tester.New(t)
	blocks, e := x.Window("s1", "a2", 1)
	assert.FatalOnError(t, e)
	assert.Integer(t, 3, len(blocks))
	assert.String(t, "a1", blocks[0].Identifier)
	assert.String(t, "a2", blocks[1].Identifier)
	assert.String(t, "edit", blocks[1].Kind)
	assert.String(t, "a3", blocks[2].Identifier)
	empty, e := x.Window("s1", "absent", 1)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, len(empty))
}

func TestSearchIndexVersionRebuild(t *testing.T) {
	d := connection.NewMemory()
	defer errors.PanicClose(d)
	path := filepath.Join(t.TempDir(), "s1.jsonl")
	system.WriteFile(path, []byte(search_tester.Transcript()), 0644)
	search_index.New(d).Append("s1", path)
	_, e := d.Exec("PRAGMA user_version = 1")
	assert.FatalOnError(t, e)
	sessions, e := search_index.New(d).Sessions()
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, len(sessions))
}

func TestSearchIndexOrdersByLatestHit(t *testing.T) {
	x, path := search_tester.New(t)
	later := filepath.Join(filepath.Dir(path), "s2.jsonl")
	system.WriteFile(
		later,
		[]byte(
			search_tester.Line(
				"b1",
				"user",
				"2026-10-08T09:00:00Z",
				false,
				`"legacy again"`,
			),
		),
		0644,
	)
	x.Append("s2", later)
	r := search_tester.Search(t, x, "legacy")
	assert.Integer(t, 2, len(r))
	assert.String(t, "s2", r[0].Session)
	assert.String(t, "s1", r[1].Session)
}
