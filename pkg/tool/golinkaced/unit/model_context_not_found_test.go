package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/base"
	"testing"
)

func TestUnknownLinkIsAnsweredQuietlyOverModelContext(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = not_found.Format("resource not found: links/%d", 12)
	assert.String(
		t,
		"resource not found: links/12",
		s.Tools().MustCallToolError(
			constant.Delete,
			map[string]any{"type": "link", "identifier": 12},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}
