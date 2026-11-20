package scope_tester

import "github.com/funtimecoding/soil/pkg/tool/goanalyze/option"

func (o *Tester) Options(
	root string,
	full bool,
	patterns ...string,
) *option.Analyze {
	o.t.Helper()
	result := option.New()
	result.Root = root
	result.Full = full
	result.Index = o.t.TempDir()
	result.Patterns = patterns

	return result
}
