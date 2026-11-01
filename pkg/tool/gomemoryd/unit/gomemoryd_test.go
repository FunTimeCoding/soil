package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/importer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"testing"
)

func TestImportFromFixtures(t *testing.T) {
	d := connection.NewMemory()
	defer errors.PanicClose(d)
	s := store.New(d)
	result, e := importer.Import(s, fixture.Path(systemConstant.MemoryPath))

	if e != nil {
		t.Fatal(e)
	}

	if result.Created != 2 {
		t.Fatalf("expected 2 created, got %d", result.Created)
	}

	result2, e := importer.Import(s, fixture.Path(systemConstant.MemoryPath))

	if e != nil {
		t.Fatal(e)
	}

	if result2.Created != 0 {
		t.Fatalf("expected 0 created on re-import, got %d", result2.Created)
	}

	if result2.Skipped != 2 {
		t.Fatalf("expected 2 skipped on re-import, got %d", result2.Skipped)
	}
}

func TestIndexEntry(t *testing.T) {
	assert.String(
		t,
		"12 cache warms on boot [build deploy] - the first request pays nothing",
		format.IndexEntry(
			&record.MemorySummary{
				Identifier:  12,
				Name:        "cache warms on boot",
				Description: "the first request pays nothing",
				Tags:        []string{"build", "deploy"},
			},
		),
	)
}

func TestIndexEntryWithoutTags(t *testing.T) {
	assert.String(
		t,
		"34 retry with backoff - double the wait after every failure",
		format.IndexEntry(
			&record.MemorySummary{
				Identifier:  34,
				Name:        "retry with backoff",
				Description: "double the wait after every failure",
			},
		),
	)
}

func TestIndexEntryWithChildren(t *testing.T) {
	assert.String(
		t,
		"56 timeouts cascade [build] - one slow upstream stalls the pool\n    + pool exhaustion, upstream budget",
		format.IndexEntry(
			&record.MemorySummary{
				Identifier:  56,
				Name:        "timeouts cascade",
				Description: "one slow upstream stalls the pool",
				Tags:        []string{"build"},
				Children:    []string{"pool exhaustion", "upstream budget"},
			},
		),
	)
}

func TestIndexEntryOmitsTimestamp(t *testing.T) {
	assert.StringNotContains(
		t,
		"2026",
		format.IndexEntry(
			&record.MemorySummary{
				Identifier:  78,
				Name:        "index rebuild is idempotent",
				Description: "running it twice changes nothing",
				UpdatedAt:   "2026-09-20T03:32:25Z",
			},
		),
	)
}

func TestAlwaysMemory(t *testing.T) {
	assert.String(
		t,
		"## cache warms on boot (12) [always]\nThe first request pays nothing.\n",
		format.AlwaysMemory(
			&record.Memory{
				Identifier: 12,
				Name:       "cache warms on boot",
				Content:    "The first request pays nothing.",
				Tags:       []string{"always"},
			},
		),
	)
}

func TestAlwaysMemoryListsChildren(t *testing.T) {
	assert.String(
		t,
		"## timeouts cascade (56) [always]\nOne slow upstream stalls the pool.\n    + pool exhaustion, upstream budget\n",
		format.AlwaysMemory(
			&record.Memory{
				Identifier: 56,
				Name:       "timeouts cascade",
				Content:    "One slow upstream stalls the pool.",
				Tags:       []string{"always"},
				Children:   []string{"pool exhaustion", "upstream budget"},
			},
		),
	)
}

func TestAlwaysMemoryOmitsDescription(t *testing.T) {
	assert.StringNotContains(
		t,
		"retrieval hook",
		format.AlwaysMemory(
			&record.Memory{
				Identifier:  12,
				Name:        "cache warms on boot",
				Content:     "The first request pays nothing.",
				Description: "a retrieval hook, not the body",
			},
		),
	)
}

func TestAlwaysMemoryPrintsItsBase(t *testing.T) {
	assert.String(
		t,
		"## error capture (12)\nBase: `../soil/doc/ai/spec/error-handling`\n- `mcp.md` - the tiers\n",
		format.AlwaysMemory(
			&record.Memory{
				Identifier: 12,
				Name:       "error capture",
				Content:    "- `mcp.md` - the tiers",
				Metadata: map[string]string{
					"base": "../soil/doc/ai/spec/error-handling",
				},
			},
		),
	)
}

func TestAlwaysMemoryPrintsEveryBase(t *testing.T) {
	assert.StringContains(
		t,
		"Base: `doc/ai/spec`, `../soil/doc/ai/spec`\n",
		format.AlwaysMemory(
			&record.Memory{
				Identifier: 12,
				Name:       "specs",
				Content:    "Read `naming.md`.",
				Metadata: map[string]string{
					"base": "doc/ai/spec, ../soil/doc/ai/spec",
				},
			},
		),
	)
}

func TestRelevantMemoryPrintsItsBase(t *testing.T) {
	assert.StringContains(
		t,
		"\nBase: `doc/ai/runbook`\nSee `fleet.md`.\n",
		format.RelevantMemory(
			&record.SearchResult{
				Identifier: 7,
				Name:       "deploy",
				Content:    "See `fleet.md`.",
				Rank:       1,
				Metadata:   map[string]string{"base": "doc/ai/runbook"},
			},
		),
	)
}

func TestSpread(t *testing.T) {
	s := token_summary.NewSpread([]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100})
	assert.Integer(t, 50, s.Median)
	assert.Integer(t, 100, s.Maximum)
	assert.Integer(t, 550, s.Total)
}

func TestSpreadIgnoresInputOrder(t *testing.T) {
	s := token_summary.NewSpread([]int{100, 10, 50, 20, 90})
	assert.Integer(t, 100, s.Maximum)
	assert.Integer(t, 50, s.Median)
}

func TestSpreadDoesNotMutateInput(t *testing.T) {
	values := []int{30, 10, 20}
	token_summary.NewSpread(values)
	assert.Integers(t, []int{30, 10, 20}, values)
}

func TestSpreadEmpty(t *testing.T) {
	s := token_summary.NewSpread(nil)
	assert.Integer(t, 0, s.Maximum)
	assert.Integer(t, 0, s.Total)
}

func TestSpreadSingleValue(t *testing.T) {
	s := token_summary.NewSpread([]int{42})
	assert.Integer(t, 42, s.Median)
	assert.Integer(t, 42, s.Maximum)
	assert.Integer(t, 42, s.Total)
}

func TestPrefixQueryAppendsToLastToken(t *testing.T) {
	assert.String(t, "scr*", web.PrefixQuery("scr"))
	assert.String(t, "memory scr*", web.PrefixQuery("memory scr"))
}

func TestPrefixQueryLeavesExistingStar(t *testing.T) {
	assert.String(t, "scr*", web.PrefixQuery("scr*"))
}

func TestPrefixQueryLeavesEmpty(t *testing.T) {
	assert.String(t, "", web.PrefixQuery(""))
	assert.String(t, "   ", web.PrefixQuery("   "))
}

func TestPrefixQueryLeavesOperatorAndPunctuation(t *testing.T) {
	assert.String(t, `foo "bar"`, web.PrefixQuery(`foo "bar"`))
	assert.String(t, "foo (", web.PrefixQuery("foo ("))
	assert.String(t, "foo:", web.PrefixQuery("foo:"))
}

func TestPrefixQueryKeepsHyphenAndDigits(t *testing.T) {
	assert.String(t, "utf-8*", web.PrefixQuery("utf-8"))
	assert.String(t, "sha256*", web.PrefixQuery("sha256"))
}
