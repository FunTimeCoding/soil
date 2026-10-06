package module_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"testing"
)

func Options(
	t *testing.T,
	root string,
	diff bool,
	full bool,
	patterns ...string,
) *option.Fix {
	t.Helper()
	result := option.New()
	result.Root = root
	result.Diff = diff
	result.Full = full
	result.Index = t.TempDir()
	result.Patterns = patterns

	return result
}
