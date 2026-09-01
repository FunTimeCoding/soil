package aleeva_report

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Parse(
	base string,
	name string,
) []*Report {
	var result []*Report
	s := system.ReadFile(base, name)

	if false {
		console.Format("Parsing: %s\n", s)
	}

	notation.MustDecode(s, &result, false)

	return result
}
