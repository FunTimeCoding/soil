package module_tester

import "github.com/funtimecoding/soil/pkg/tool/gofix/option"

func (o *Tester) Options(
	root string,
	diff bool,
	full bool,
	patterns ...string,
) *option.Fix {
	o.t.Helper()
	result := option.New()
	result.Root = root
	result.Diff = diff
	result.Full = full
	result.Index = o.t.TempDir()
	result.Patterns = patterns

	return result
}
