package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service/format"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"testing"
)

func TestIndexEntry(t *testing.T) {
	assert.String(
		t,
		"12 cache warms on boot [build deploy] - the first request pays nothing",
		format.IndexEntry(
			&store.MemorySummary{
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
			&store.MemorySummary{
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
			&store.MemorySummary{
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
			&store.MemorySummary{
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
			&store.Memory{
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
			&store.Memory{
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
			&store.Memory{
				Identifier:  12,
				Name:        "cache warms on boot",
				Content:     "The first request pays nothing.",
				Description: "a retrieval hook, not the body",
			},
		),
	)
}
