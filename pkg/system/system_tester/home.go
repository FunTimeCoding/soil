package system_tester

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func Home(t *testing.T) string {
	t.Helper()
	result := t.TempDir()
	t.Setenv(soil.HomeEnvironment, result)

	return result
}
