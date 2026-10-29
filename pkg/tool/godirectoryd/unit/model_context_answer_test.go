package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"testing"
)

func TestAQuietClassIsAnsweredWithoutACaptureOverModelContext(t *testing.T) {
	m := memory.New()
	service_tester.Connect(t, service_tester.NewSource(), m).MustCallToolError(
		constant.DeleteUser,
		map[string]any{"account": "ghost"},
	)
	assert.Count(t, 0, m.Events())
}

func TestAnUnclassifiedFailureIsCapturedOverModelContext(t *testing.T) {
	c := service_tester.NewSource()
	c.Failure = errors.New("directory answered nonsense")
	m := memory.New()
	service_tester.Connect(t, c, m).MustCallToolError(
		constant.ListUser,
		map[string]any{},
	)
	assert.Count(t, 1, m.Events())
}
