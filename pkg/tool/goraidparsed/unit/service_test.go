package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/unit/service_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestGenerateAnswersAMissingInput(t *testing.T) {
	_, e := service_tester.New(t).Generate(
		[]string{filepath.Join(t.TempDir(), "absent.zevtc")},
		time.Now(),
	)
	assert.True(t, not_found.Is(e))
}
