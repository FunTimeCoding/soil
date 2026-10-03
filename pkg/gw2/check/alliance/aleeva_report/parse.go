package aleeva_report

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Parse(
	base string,
	name string,
) []*Report {
	var result []*Report
	notation.MustDecode(system.ReadFile(base, name), &result, false)

	return result
}
