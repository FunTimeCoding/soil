package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/hypertext"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestParsedFixturesLeaveNoHandleOpen(t *testing.T) {
	before := fixture.OpenHandles(t)
	t.Run(
		"parse",
		func(t *testing.T) {
			for range 100 {
				assert.NotNil(
					t,
					hypertext.Document(
						openFixture(
							t,
							systemConstant.HypertextPath,
							"test.html",
						),
					),
				)
			}
		},
	)
	assert.True(t, fixture.OpenHandles(t)-before < 20)
}
