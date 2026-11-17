package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/base"
	"testing"
)

func TestQuietClassIsAnsweredWithoutCaptureOverModelContext(t *testing.T) {
	s := base.New(t)
	s.Tools().MustCallToolError(
		constant.DeleteUser,
		map[string]any{"account": "ghost"},
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUnclassifiedFailureIsCapturedOverModelContext(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = errors.New("directory answered nonsense")
	s.Tools().MustCallToolError(constant.ListUser, map[string]any{})
	assert.Count(t, 1, s.Reporter.Events())
}
