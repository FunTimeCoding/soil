package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"testing"
)

func TestCensusLinesSortedByReasonThenPath(t *testing.T) {
	var r output.Results
	r.AddUnchecked("doc/b.md", 3, "~/x", constant.ReasonHome)
	r.AddUnchecked("doc/a.md", 9, "/proc/", constant.ReasonSystem)
	r.AddUnchecked("doc/a.md", 2, "~/y", constant.ReasonHome)
	assert.Strings(
		t,
		[]string{
			"doc/a.md:2: unchecked home ~/y",
			"doc/b.md:3: unchecked home ~/x",
			"doc/a.md:9: unchecked system /proc/",
			"unchecked home: 2",
			"unchecked system: 1",
			"unchecked: 3",
		},
		output.CensusLines(r.Unchecked),
	)
}

func TestCensusLinesEmpty(t *testing.T) {
	assert.Strings(t, nil, output.CensusLines(nil))
}
