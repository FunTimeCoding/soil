package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"path/filepath"
	"testing"
)

func NewOption(
	t *testing.T,
	s *mock_system.System,
) *option.Flight {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "goflightd.db")
	result.Command = s.Run
	result.Stream = s.Stream

	return result
}
