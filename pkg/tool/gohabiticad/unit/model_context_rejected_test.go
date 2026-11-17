package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/base"
	"testing"
)

func TestInvalidStatIsAnsweredQuietlyOverModelContext(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = validation.New("%q is not a valid Stat.", "bravo")
	assert.String(
		t,
		`"bravo" is not a valid Stat.`,
		s.Tools().MustCallToolError(
			constant.AllocateStat,
			map[string]any{constant.Stat: "bravo"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}
