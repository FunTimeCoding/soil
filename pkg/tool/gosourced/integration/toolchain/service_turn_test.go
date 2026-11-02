package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"testing"
)

func TestTwoSessionsRenameInOneModuleWithoutRefusing(t *testing.T) {
	d := renameFollowFixture(t)
	s := service_tester.NewService(t)
	results := make(chan *output.Results, 2)
	failures := make(chan error, 2)

	for _, rename := range [][3]string{
		{"alfa", "OldName", "NewName"},
		{"bravo", "OldType", "NewType"},
	} {
		go func() {
			r, e := s.Rename(
				turn.WithSession(t.Context(), rename[0]),
				d,
				"example/pkg/alfa",
				rename[1],
				rename[2],
				"",
				false,
			)
			failures <- e
			results <- r
		}()
	}

	for range 2 {
		assert.FatalOnError(t, <-failures)
		testutil.AssertBlocked(t, <-results, 0)
	}

	assert.StringContains(
		t,
		"func NewName()",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/new_name.go"),
	)
	assert.StringContains(
		t,
		"type NewType",
		service_tester.ReadFixtureFile(t, d, "pkg/alfa/new_type.go"),
	)
}
