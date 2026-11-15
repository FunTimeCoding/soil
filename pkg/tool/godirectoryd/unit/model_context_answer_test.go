package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"testing"
)

func TestQuietClassIsAnsweredWithoutCaptureOverModelContext(t *testing.T) {
	m := memory.New()
	base.NewWith(t, service_tester.NewSource(), m).Connect(t).MustCallToolError(
		constant.DeleteUser,
		map[string]any{"account": "ghost"},
	)
	assert.Count(t, 0, m.Events())
}

func TestUnclassifiedFailureIsCapturedOverModelContext(t *testing.T) {
	c := service_tester.NewSource()
	c.Failure = errors.New("directory answered nonsense")
	m := memory.New()
	base.NewWith(t, c, m).Connect(t).MustCallToolError(
		constant.ListUser,
		map[string]any{},
	)
	assert.Count(t, 1, m.Events())
}
