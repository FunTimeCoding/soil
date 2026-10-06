package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"testing"
)

func TestSearchHybridFusesKeywordAndVectorResults(t *testing.T) {
	s := store_tester.OpenTestStore(t)
	defer s.Close()
	d := t.TempDir()
	store_tester.WriteFixture(t, d, "alfa.md", "# Alfa\n\nzulu appears here.\n")
	store_tester.WriteFixture(
		t,
		d,
		"bravo.md",
		"# Bravo\n\nnothing to match.\n",
	)
	s.AddCollection("test", d, constant.DefaultGlob)
	s.Index("test")
	now := "2024-01-01T00:00:00Z"
	s.InsertEmbedding(
		store.HashContent("# Alfa\n\nzulu appears here.\n"),
		0,
		0,
		[]float32{0, 1},
		now,
	)
	s.InsertEmbedding(
		store.HashContent("# Bravo\n\nnothing to match.\n"),
		0,
		0,
		[]float32{1, 0},
		now,
	)
	results, e := s.SearchHybrid(
		search_option.New("zulu", 10),
		new(fixedEmbedder),
	)
	assert.FatalOnError(t, e)
	var paths []string

	for _, r := range results {
		paths = append(paths, r.Path)
	}

	assert.Strings(t, []string{"alfa.md", "bravo.md"}, paths)
}
