package aleeva_report

import (
	"fmt"
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
		fmt.Printf("Parsing: %s\n", s)
	}

	notation.MustDecode(s, &result, false)

	return result
}
