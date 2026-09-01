package result

import "github.com/funtimecoding/soil/pkg/console"

func (r *Result) Print() {
	if r.Exit != 0 || r.Error != nil {
		console.Format("Error (%d): %v\n", r.Exit, r.Error)
	}

	if r.OutputString != "" {
		console.Line("Stdout:")
		console.Line(r.OutputString)
	}

	if r.ErrorString != "" {
		console.Line("Stderr:")
		console.Line(r.ErrorString)
	}
}
