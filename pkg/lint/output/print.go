package output

import "github.com/funtimecoding/soil/pkg/console"

func Print(
	r *Results,
	summary bool,
) bool {
	if r.Header != "" {
		console.Line(r.Header)
	}

	return PrintResults(r.Entries, summary)
}
