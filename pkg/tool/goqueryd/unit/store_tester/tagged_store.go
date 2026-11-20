package store_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (o *Tester) TaggedStore() *store.Store {
	o.t.Helper()
	s := o.OpenTestStore()
	directory := o.t.TempDir()
	o.WriteFixture(directory, "alfa.md", "# Alfa\n\nKeyword quasar.\n")
	o.WriteFixture(directory, "bravo.md", "# Bravo\n\nKeyword quasar.\n")
	o.WriteFixture(directory, "charlie.md", "# Charlie\n\nKeyword quasar.\n")
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.Index("test")
	assert.FatalOnError(
		o.t,
		s.SetMetadata(
			"test",
			"alfa.md",
			map[string][]string{
				constant.FixtureTagKey:    {constant.FixtureBuildValue, "groom"},
				constant.FixtureAuthorKey: {"one"},
			},
		),
	)
	assert.FatalOnError(
		o.t,
		s.SetMetadata(
			"test",
			"bravo.md",
			map[string][]string{
				constant.FixtureTagKey:    {constant.FixtureBuildValue},
				constant.FixtureAuthorKey: {"two"},
			},
		),
	)
	assert.FatalOnError(
		o.t,
		s.SetMetadata(
			"test",
			"charlie.md",
			map[string][]string{
				constant.FixtureTagKey:    {"groom"},
				constant.FixtureAuthorKey: {"one"},
			},
		),
	)

	return s
}
