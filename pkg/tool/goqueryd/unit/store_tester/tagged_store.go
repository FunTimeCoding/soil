package store_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func TaggedStore(t *testing.T) *store.Store {
	t.Helper()
	s := OpenTestStore(t)
	directory := t.TempDir()
	WriteFixture(t, directory, "alfa.md", "# Alfa\n\nKeyword quasar.\n")
	WriteFixture(t, directory, "bravo.md", "# Bravo\n\nKeyword quasar.\n")
	WriteFixture(t, directory, "charlie.md", "# Charlie\n\nKeyword quasar.\n")
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.Index("test")
	s.SetMetadata(
		"test",
		"alfa.md",
		map[string][]string{
			constant.FixtureTagKey:    {constant.FixtureBuildValue, "groom"},
			constant.FixtureAuthorKey: {"one"},
		},
	)
	s.SetMetadata(
		"test",
		"bravo.md",
		map[string][]string{
			constant.FixtureTagKey:    {constant.FixtureBuildValue},
			constant.FixtureAuthorKey: {"two"},
		},
	)
	s.SetMetadata(
		"test",
		"charlie.md",
		map[string][]string{
			constant.FixtureTagKey:    {"groom"},
			constant.FixtureAuthorKey: {"one"},
		},
	)

	return s
}
