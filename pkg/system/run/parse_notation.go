package run

import "github.com/funtimecoding/soil/pkg/notation"

func (r *Run) ParseNotation(a any) {
	notation.MustDecode(r.OutputString, &a, r.Verbose)
}
