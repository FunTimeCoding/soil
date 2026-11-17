package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"testing"
)

func TestModelContextAnswersClassifiedRefusalQuietly(t *testing.T) {
	s := base.New(t)
	s.Alfa().FailProjects(not_found.New("group", "kilo"))
	assert.StringContains(
		t,
		"group not found: kilo",
		s.Tools().MustCallToolError(constant.ListProjects, nil),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestModelContextCapturesUnclassifiedFailure(t *testing.T) {
	s := base.New(t)
	s.Alfa().FailProjects(errors.New("refused for no known reason"))
	assert.StringContains(
		t,
		"refused for no known reason",
		s.Tools().MustCallToolError(constant.ListProjects, nil),
	)
	assert.Count(t, 1, s.Reporter.Events())
}
