package scope_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"testing"
)

func Options(
	t *testing.T,
	root string,
	full bool,
	patterns ...string,
) *option.Analyze {
	t.Helper()
	result := option.New()
	result.Root = root
	result.Full = full
	result.Index = t.TempDir()
	result.Patterns = patterns

	return result
}
