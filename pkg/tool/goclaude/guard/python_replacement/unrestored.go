package python_replacement

import "github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"

func Unrestored(s *shell_script.Script) bool {
	var replacements []*Replacement
	var calls []*shell_script.Call

	for _, c := range s.Calls {
		if !c.TopLevel {
			continue
		}

		calls = append(calls, c)
		code, extracted := source(c)

		if !extracted {
			continue
		}

		if file, found := single(code); found {
			replacements = append(
				replacements,
				New(c.Statement.End().Offset(), file),
			)
		}
	}

	for _, r := range replacements {
		if !restored(r, calls) {
			return true
		}
	}

	return false
}
