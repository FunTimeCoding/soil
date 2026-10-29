package base

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
	"path/filepath"
	"testing"
)

func NewOption(t *testing.T) *option.Telemetry {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "gotelemetryd.db")

	return result
}
